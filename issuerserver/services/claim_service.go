package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-merkletree-sql"
	mtsql "github.com/iden3/go-merkletree-sql/db/sql"
	"github.com/juju/errors"
	common2 "issuerserver/common"
	"issuerserver/config"
	"issuerserver/models/common"
	"issuerserver/models/database"
	"issuerserver/pkg/blockchain"
	"issuerserver/pkg/codec"
	"issuerserver/pkg/db"
	"issuerserver/pkg/logging"
	"issuerserver/store"
	"log"
	"math/big"
	"math/rand"
	"time"
)

func IssueClaim(claimRequest *common.ClaimRequest) (*core.Claim, error) {
	log.Println("start adding claim")
	ctx := context.Background()
	//Build claim
	claim, err := generateClaimInfo(claimRequest)
	if err != nil {
		return nil, err
	}
	logging.Info.Println("claim is built")
	//Add this claim to mysql claim table .(this does not modify trees which is performed later)
	saveClaimBody(claim, claimRequest.IssuerDid)
	logging.Info.Println("claim is saved")
	//Find issuer from database
	issuer, err := store.SelectDidById(claimRequest.IssuerDid)
	if err != nil {
		log.Printf("select did failed: %s", claimRequest.IssuerDid)
		return nil, err
	}
	logging.Info.Printf("issuer loaded:%s", issuer.Did)
	//Load trees
	cltTree, rotTree, torTree, err := loadTrees(ctx, issuer)
	logging.Info.Println("Tree loaded")
	fmt.Println("claims root:")
	fmt.Println(cltTree.Root().Hex())
	oldCltRoot, oldRotRoot, oldTorRoot := cltTree.Root(), rotTree.Root(), torTree.Root()
	oldIdentityState, _ := merkletree.HashElems(oldCltRoot.BigInt(), oldRotRoot.BigInt(), oldTorRoot.BigInt())

	dbg, _ := core.IDFromString(issuer.Did)
	fmt.Println("dbg--before input")
	fmt.Println(oldCltRoot.BigInt(), oldRotRoot.BigInt(), oldTorRoot.BigInt())
	fmt.Println(hexutil.Encode(dbg.Bytes()))
	fmt.Println(oldIdentityState.Hex())

	logging.Info.Printf("Old identity state:%s", oldIdentityState.String())
	//Load auth claim
	authClaimHIndex := new(big.Int).SetBytes(issuer.AuthClaimHi)
	//Add this claim to merkle tree (which is stored in postgresql)
	hIndex, hValue, _ := claim.HiHv()
	err = cltTree.Add(ctx, hIndex, hValue)
	if err != nil {
		return nil, err
	}
	logging.Info.Println("Adding claim to claims tree")
	//Recalculate state root
	newIdentityState, err := merkletree.HashElems(
		cltTree.Root().BigInt(),
		rotTree.Root().BigInt(),
		torTree.Root().BigInt(),
	)
	if err != nil {
		return nil, err
	}
	logging.Info.Printf("New identity state %s", newIdentityState.String())
	//Create state transition proof: move from old state to new state
	authClaim, _, err := loadClaim(authClaimHIndex)
	authClaimExistProof, _, _ := cltTree.GenerateProof(ctx, authClaimHIndex, oldCltRoot)
	authClaimNonRevProof, _, _ := rotTree.GenerateProof(ctx, new(big.Int).SetUint64(authClaim.GetRevocationNonce()), oldRotRoot)
	stateTransitionInputs, err := createStateTransitionInput(issuer, oldIdentityState, oldCltRoot, oldRotRoot, oldTorRoot, newIdentityState, authClaim, authClaimExistProof, authClaimNonRevProof)
	if err != nil {
		fmt.Println("failed to create state transition inputs")
		return nil, err
	}
	inputBytes, err := stateTransitionInputs.InputsMarshal()
	if err != nil {
		return nil, err
	}
	logging.Info.Println("state transition input created complete:")
	fmt.Println(string(inputBytes))
	//Sync state to blockchain
	err = blockchain.SyncIdentityState(issuer.Did, stateTransitionInputs)
	if err != nil {
		log.Printf("Failed to sync did %s to blockchain new state %s", issuer.Did, newIdentityState)
		return nil, err
	}
	logging.Info.Println("Blockchain sync complete")
	//Update did with synced
	issuer.IsGenesis = false
	issuer.UpdatedAt = time.Now()
	store.UpdateDid(issuer)
	//Update offer claimer
	logging.Info.Println("issuer state updated complete")
	return claim, nil
}

func RevokeClaim(claimHi *big.Int) error {
	log.Printf("start revoke claim %s", claimHi)
	ctx := context.Background()
	//Read claim from t_claim and issuer
	claim, claimdDb, issuer, err := loadIssuerAndClaim(claimHi, true)
	if err != nil {
		fmt.Println("failed to call loadIssuerAndClaim ")
		return err
	}
	//Load trees
	cltTree, rotTree, torTree, err := loadTrees(ctx, issuer)
	oldCltRoot, oldRotRoot, oldTorRoot := cltTree.Root(), rotTree.Root(), torTree.Root()
	oldIdentityState, _ := merkletree.HashElems(oldCltRoot.BigInt(), oldRotRoot.BigInt(), oldTorRoot.BigInt())
	//Load auth claim
	authClaimHIndex := new(big.Int).SetBytes(issuer.AuthClaimHi)
	//Add revocation nonce as index to leaf and update root
	err = rotTree.Add(ctx, big.NewInt(int64(claim.GetRevocationNonce())), big.NewInt(int64(claim.GetVersion())))
	if err != nil {
		log.Printf("Failed to add revocation nonce %d to revocaiton tree", claim.GetRevocationNonce())
		return err
	}
	newIdentityState, _ := merkletree.HashElems(
		cltTree.Root().BigInt(),
		rotTree.Root().BigInt(),
		torTree.Root().BigInt())
	//Create state transition proof: move from old state to new state
	authClaim, _, err := loadClaim(authClaimHIndex)
	authClaimExistProof, _, _ := cltTree.GenerateProof(ctx, authClaimHIndex, oldCltRoot)
	authClaimNonRevProof, _, _ := rotTree.GenerateProof(ctx, new(big.Int).SetUint64(authClaim.GetRevocationNonce()), oldRotRoot)
	stateTransitionInputs, err := createStateTransitionInput(issuer, oldIdentityState, oldCltRoot, oldRotRoot, oldTorRoot, newIdentityState, authClaim, authClaimExistProof, authClaimNonRevProof)
	if err != nil {
		fmt.Println("failed to create state transition inputs")
		return err
	}
	//Sync state to blockchain
	err = blockchain.SyncIdentityState(issuer.Did, stateTransitionInputs)
	if err != nil {
		log.Printf("Failed to sync did %s to blockchain new state %s", issuer.Did, newIdentityState)
		return err
	}
	//Update did with synced
	issuer.IsGenesis = false
	issuer.UpdatedAt = time.Now()
	err = store.UpdateDid(issuer)
	if err != nil {
		return errors.Trace(err)
	}
	//Update revoke status
	err = store.UpdateClaimRevokeStatus(claimdDb.Id, true)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// Get
func GetNonRevocationProofForClaim(did string, claimHi *big.Int) (*merkletree.Proof, error) {
	//Read claim from t_claim and issuer
	claim, _, err := loadClaim(claimHi)
	if err != nil {
		fmt.Println("failed to call GetNonRevocationProofForClaim ")
		return nil, err
	}

	revocationNonce := claim.GetRevocationNonce()
	//Load did info from db
	didData, err := store.SelectDidById(did)
	if err != nil {
		log.Printf("Failed to select did %s", did)
		return nil, err
	}
	//Load revocation tree
	ctx := context.Background()
	dbStorage := mtsql.NewSqlStorage(db.GetPgDb(), didData.RotId)
	rotTree, err := merkletree.NewMerkleTree(ctx, dbStorage, 32)
	if err != nil {
		log.Printf("Failed to load rot tree %d", didData.RotId)
		return nil, err
	}
	//Generate proof.
	proof, _, err := rotTree.GenerateProof(ctx, new(big.Int).SetUint64(revocationNonce), rotTree.Root())
	if err != nil {
		fmt.Println("Failed to generate proof ")
		return nil, err
	}
	return proof, nil
}

func generateClaimInfo(claimInfo *common.ClaimRequest) (*core.Claim, error) {
	//Build claim
	schema, _ := core.NewSchemaHashFromHex(config.GetConfig().Schema.DefaultHash)
	revocationNonce := rand.Uint64()
	subjectId, _ := core.IDFromString(claimInfo.HolderDid)
	claim, _ := core.NewClaim(schema,
		core.WithExpirationDate(claimInfo.ExpiredAt),
		core.WithRevocationNonce(revocationNonce),
		core.WithIndexID(subjectId),
	)

	//Encode dcp claim data
	dcpData := &common.DCP{
		DataCategory: claimInfo.DataCategory,
		SubCategory:  claimInfo.SubCategory,
		LowerBoundary: common.Boundary{
			Include: claimInfo.LowerBoundInclude,
			Value:   claimInfo.LowerBoundValue,
		},
		UpperBoundary: common.Boundary{
			Include: claimInfo.UpperBoundInclude,
			Value:   claimInfo.UpperBoundValue,
		},
		HolderAddress: claimInfo.HolderAddress,
	}
	err := codec.EncodeDcpClaim(claim, dcpData)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return claim, nil
}

func loadClaim(claimHi *big.Int) (*core.Claim, *database.ClaimData, error) {
	claimDb, err := store.SelectClaimByHIndex(claimHi)
	if err != nil {
		log.Printf("Failed to load claim %s", claimHi)
		return nil, nil, err
	}
	claim := &core.Claim{}
	err = claim.UnmarshalBinary(claimDb.ClaimBinary)
	if err != nil {
		log.Printf("Failed to unmarshal bytes")
		return nil, nil, err
	}
	return claim, claimDb, nil
}

func loadIssuerAndClaim(claimHi *big.Int, forRevoke bool) (*core.Claim, *database.ClaimData, *database.DidData, error) {
	//Load claim
	claim, claimDb, err := loadClaim(claimHi)
	if forRevoke && claimDb.Revoked {
		return nil, nil, nil, errors.New("Claim already revoked")
	}
	//Load did
	issuerDid := claimDb.Issuer
	issuer, err := store.SelectDidById(issuerDid)
	if err != nil {
		log.Printf("select did failed: %s", issuerDid)
		return nil, nil, nil, err
	}
	return claim, claimDb, issuer, nil
}

func loadTrees(ctx context.Context, identity *database.DidData) (*merkletree.MerkleTree, *merkletree.MerkleTree, *merkletree.MerkleTree, error) {
	log.Printf("start loading trees for %s", identity.Did)
	cltTree, err := loadTree(ctx, identity.CltId)
	if err != nil {
		return nil, nil, nil, err
	}
	rotTree, err := loadTree(ctx, identity.RotId)
	if err != nil {
		return nil, nil, nil, err
	}
	torTree, err := loadTree(ctx, identity.TorId)
	if err != nil {
		return nil, nil, nil, err
	}
	return cltTree, rotTree, torTree, nil
}

func loadTree(ctx context.Context, mtId uint64) (*merkletree.MerkleTree, error) {
	tree, err := merkletree.NewMerkleTree(ctx, mtsql.NewSqlStorage(db.GetPgDb(), mtId), 32)
	if err != nil {
		log.Printf("Failed to reload tree :%d", mtId)
		return nil, err
	}
	log.Println("load tree with root: ", tree.Root().Hex())
	return tree, nil
}

func createStateTransitionInput(identity *database.DidData, oldState, oldCltRoot, oldRotRoot, oldTorRoot, newState *merkletree.Hash,
	authClaim *core.Claim, authClaimExistProof, authClaimNonRevProof *merkletree.Proof) (*circuits.StateTransitionInputs, error) {
	didObj, err := core.IDFromString(identity.Did)
	if err != nil {
		fmt.Printf("Failed to get core.Did from str:%s", identity.Did)
		return nil, err
	}
	hashOldAndNewStates, err := poseidon.Hash([]*big.Int{oldState.BigInt(), newState.BigInt()})
	if err != nil {
		log.Println("Failed to get hashOldAndNewStates")
	}
	var pkeySlice [32]byte
	copy(pkeySlice[:], identity.AuthPrivateKey)
	privateKey := babyjub.PrivateKey(pkeySlice)

	signature := privateKey.SignPoseidon(hashOldAndNewStates)
	stateTransitionInputs := &circuits.StateTransitionInputs{
		ID: &didObj,
		OldTreeState: circuits.TreeState{
			State:          oldState,
			ClaimsRoot:     oldCltRoot,
			RevocationRoot: oldRotRoot,
			RootOfRoots:    oldTorRoot,
		},
		NewState:          newState,
		IsOldStateGenesis: identity.IsGenesis,
		AuthClaim: circuits.Claim{
			Claim: authClaim,
			Proof: authClaimExistProof,
			NonRevProof: &circuits.ClaimNonRevStatus{
				Proof: authClaimNonRevProof,
			},
		},
		Signature: signature,
	}
	return stateTransitionInputs, nil
}

func saveClaimBody(claim *core.Claim, issuer string) error {
	//Extract meta
	hIndex, err := claim.HIndex()
	if err != nil {
		return err
	}
	binary, err := claim.MarshalBinary()
	if err != nil {
		return err
	}

	//Build
	claimDbModel := database.ClaimData{
		ClaimHi:     hIndex.Bytes(),
		ClaimBinary: binary,
		Issuer:      issuer,
		Revoked:     false,
		ClaimType:   common2.CLAIM_TYPE_DCP,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = store.InsertClaim(&claimDbModel)
	if err != nil {
		return err
	}
	return nil
}

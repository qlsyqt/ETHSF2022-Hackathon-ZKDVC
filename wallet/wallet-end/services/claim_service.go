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
	"log"
	"math/big"
	"time"
	common2 "wallet-end/common"
	"wallet-end/models/database"
	"wallet-end/models/response"
	"wallet-end/pkgs/blockchain"
	"wallet-end/pkgs/codec"
	"wallet-end/store"
)

func FetchClaims() ([]*response.FetchClaimVO, error) {
	claimsDb, err := store.SelectClaimsByBatch(-1, 99999)
	if err != nil {
		return nil, err
	}
	ret := make([]*response.FetchClaimVO, 0)
	for _, c := range claimsDb {
		if c.ClaimType != "dcp" {
			continue
		}
		claim := &core.Claim{}
		err = claim.UnmarshalBinary(c.ClaimBinary)
		if err != nil {
			return nil, err
		}
		dcp := codec.DecodeDcpClaim(claim)
		vo := response.FetchClaimVO{
			HIndex:       hexutil.Encode(c.ClaimHi),
			DataCategory: dcp.DataCategory,
			SubCategory:  dcp.SubCategory,
			Interval:     codec.RangeToString(&dcp.LowerBoundary, &dcp.UpperBoundary),
		}
		ret = append(ret, &vo)
	}

	return ret, nil
}

func DoAddClaim(claim *core.Claim, issuer string) (*core.Claim, error) {
	log.Println("start adding claim")
	ctx := context.Background()
	//Add this claim to mysql claim table .(this does not modify trees which is performed later)
	saveClaimBody(claim, issuer)
	log.Println("claim is saved")
	//Find issuer from database
	holderData, err := store.SelectDidById(myDid.String())
	if err != nil {
		log.Printf("select holder did failed: %s", myDid.String())
		return nil, err
	}
	log.Printf("holder loaded from db:%s", myDid.String())
	cltTree, rotTree, torTree, err := loadTrees(ctx, holderData)
	log.Println("Tree loaded")
	fmt.Println("claims root:")
	fmt.Println(cltTree.Root().Hex())
	oldCltRoot, oldRotRoot, oldTorRoot := cltTree.Root(), rotTree.Root(), torTree.Root()
	oldIdentityState, _ := merkletree.HashElems(oldCltRoot.BigInt(), oldRotRoot.BigInt(), oldTorRoot.BigInt())

	dbg, _ := core.IDFromString(holderData.Did)
	fmt.Println(hexutil.Encode(dbg.Bytes()))
	fmt.Println(hexutil.Encode(oldIdentityState.BigInt().Bytes()))

	log.Printf("Old identity state:%s", oldIdentityState.String())
	//Load auth claim
	authClaimHIndex := new(big.Int).SetBytes(holderData.AuthClaimHi)
	//Add this claim to merkle tree (which is stored in postgresql)
	hIndex, hValue, _ := claim.HiHv()
	err = cltTree.Add(ctx, hIndex, hValue)
	if err != nil {
		return nil, err
	}
	log.Println("Adding claim to claims tree")
	//Recalculate state root
	newIdentityState, err := merkletree.HashElems(
		cltTree.Root().BigInt(),
		rotTree.Root().BigInt(),
		torTree.Root().BigInt(),
	)
	if err != nil {
		return nil, err
	}
	log.Printf("New identity state %s", newIdentityState.String())
	//Create state transition proof: move from old state to new state
	authClaim, _, err := loadClaim(authClaimHIndex)
	authClaimExistProof, _, _ := cltTree.GenerateProof(ctx, authClaimHIndex, oldCltRoot)
	authClaimNonRevProof, _, _ := rotTree.GenerateProof(ctx, new(big.Int).SetUint64(authClaim.GetRevocationNonce()), oldRotRoot)
	stateTransitionInputs, err := createStateTransitionInput(holderData, oldIdentityState, oldCltRoot, oldRotRoot, oldTorRoot, newIdentityState, authClaim, authClaimExistProof, authClaimNonRevProof)
	if err != nil {
		fmt.Println("failed to create state transition inputs")
		return nil, err
	}
	inputBytes, err := stateTransitionInputs.InputsMarshal()
	if err != nil {
		return nil, err
	}
	log.Println("state transition input created complete:")
	fmt.Println(string(inputBytes))
	//Sync state to blockchain
	err = blockchain.SyncIdentityState(holderData.Did, stateTransitionInputs)
	if err != nil {
		log.Printf("Failed to sync did %s to blockchain new state %s", holderData.Did, newIdentityState)
		return nil, err
	}
	log.Println("Blockchain sync complete")
	//Update did with synced
	holderData.IsGenesis = false
	holderData.UpdatedAt = time.Now()
	store.UpdateDid(holderData)
	log.Println("holderData state updated complete")
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

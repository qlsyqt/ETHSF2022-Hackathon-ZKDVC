package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql/db/sql"
	"github.com/juju/errors"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
	"wallet-end/common"
	"wallet-end/config"
	"wallet-end/models/database"
	"wallet-end/pkgs/db"
	"wallet-end/store"
	"wallet-end/utils"

	//"github.com/iden3/go-iden3-crypto/utils"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
	"sync"
)

var lck sync.Mutex
var myDid *core.ID

func FetchIdentity() (string, error) {
	cfg := config.GetConfig()
	lck.Lock()
	defer lck.Unlock()

	if myDid != nil {
		return myDid.String(), nil
	}

	did, err := deduceIdentity(cfg.PolygonId.PrivateKey, cfg.PolygonId.AuthRevocationNonce)
	if err != nil {
		return "", err
	}
	exist, err := identityExisted(did.String())
	fmt.Println("identity existed ", err)
	if err != nil {
		return "", err
	}
	if !exist {
		log.Println("did not existed yet. Create one")
		createIdentityByPrivateKey(cfg.PolygonId.PrivateKey, cfg.PolygonId.AuthRevocationNonce)
	}
	fmt.Println("create complete")
	myDid = did

	return did.String(), nil
}

func identityExisted(didStr string) (bool, error) {
	result, err := store.SelectDidById(didStr)
	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	fmt.Println(result.Did)
	return true, nil
}

func deduceIdentity(privateKey string, revNonce uint64) (*core.ID, error) {
	_, publicKey, err := utils.HexPrivateKeyToBjjKeypair(privateKey)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//create auth claim
	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(publicKey.X, publicKey.Y),
		core.WithRevocationNonce(revNonce))
	//create trees
	ctx := context.Background()
	clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	ror, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	hIndex, hValue, _ := authClaim.HiHv()
	clt.Add(ctx, hIndex, hValue)
	state, _ := merkletree.HashElems(
		clt.Root().BigInt(),
		rot.Root().BigInt(),
		ror.Root().BigInt(),
	)

	id, _ := core.IdGenesisFromIdenState(core.TypeDefault, state.BigInt())
	return id, nil
}

func createIdentityByPrivateKey(privateKeyHex string, revNonce uint64) (string, error) {
	//1. Create auth claim as the very first claim under the identity
	babyJubjubPrivKey, babyJubjubPubKey, err := utils.HexPrivateKeyToBjjKeypair(privateKeyHex)
	if err != nil {
		return "", errors.Trace(err)
	}

	//create auth claim
	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(babyJubjubPubKey.X, babyJubjubPubKey.Y),
		core.WithRevocationNonce(revNonce))

	ctx := context.Background()
	authClaimHi, _ := authClaim.HIndex()
	authClaimBinary, _ := authClaim.MarshalBinary()
	// Generate claims tree
	did, cltId, rotId, torId, _, err := initAndSaveIdentityTrees(ctx, authClaim)
	if err != nil {
		log.Println("Identity create failed")
		return "", err
	}
	// Save auth claim to mysql
	authClaimObj := database.ClaimData{
		ClaimHi:     authClaimHi.Bytes(),
		ClaimBinary: authClaimBinary,
		Issuer:      "",
		Revoked:     false,
		ClaimType:   common.CLAIM_TYPE_AUTH,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = store.InsertClaim(&authClaimObj)
	if err != nil {
		log.Println("Fail to insert claim body to db")
		return "", err
	}
	// Save to identity table
	didData := &database.DidData{
		Did:            did,
		AuthPrivateKey: babyJubjubPrivKey[:],
		AuthClaimHi:    authClaimHi.Bytes(),
		IsGenesis:      true,
		CltId:          cltId,
		RotId:          rotId,
		TorId:          torId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err = store.InsertDid(didData)
	if err != nil {
		log.Println("Identity insert failed")
		return "", err
	}
	fmt.Println("insert did complete")
	return did, nil
}

func initAndSaveIdentityTrees(ctx context.Context, authClaim *core.Claim) (string, uint64, uint64, uint64, *merkletree.Hash, error) {
	// Create empty Claims tree
	cltId := rand.Uint64() & 0x7111111111111111 //uint64 with highest bit set to 1 cannot be saved by datbase/sql
	cltStore := sql.NewSqlStorage(db.GetPgDb(), cltId)
	clt, err := merkletree.NewMerkleTree(ctx, cltStore, 32)
	if err != nil {
		log.Println("Claim tree create failed")
		fmt.Println(err)
		return "", 0, 0, 0, nil, err
	}
	// Create empty Revocation tree
	rotId := rand.Uint64() & 0x7111111111111111 //uint64 with highest bit set to 1 cannot be saved by datbase/sql
	rotStore := sql.NewSqlStorage(db.GetPgDb(), rotId)
	rot, err := merkletree.NewMerkleTree(ctx, rotStore, 32)
	if err != nil {
		log.Println("Revocation tree create failed")
		return "", 0, 0, 0, nil, err
	}
	// Create empty Roots tree
	torId := rand.Uint64() & 0x7111111111111111 //uint64 with highest bit set to 1 cannot be saved by datbase/sql
	torStore := sql.NewSqlStorage(db.GetPgDb(), torId)
	tor, err := merkletree.NewMerkleTree(ctx, torStore, 32)
	if err != nil {
		log.Println("Roots tree create failed")
		return "", 0, 0, 0, nil, err
	}

	//Add auth claim to claims
	hIndex, hValue, err := authClaim.HiHv()
	if err != nil {
		log.Println("Auth claim hash failed")
		return "", 0, 0, 0, nil, err
	}
	clt.Add(ctx, hIndex, hValue)
	state, err := merkletree.HashElems(
		clt.Root().BigInt(),
		rot.Root().BigInt(),
		tor.Root().BigInt())
	fmt.Println("dbg-fetchIdentity")
	fmt.Println(state.Hex())
	fmt.Println(clt.Root().BigInt())
	fmt.Println(rot.Root().BigInt())
	fmt.Println(rot.Root().BigInt())
	if err != nil {
		log.Println("Genesis state failed failed")
		return "", 0, 0, 0, nil, err
	}

	did, err := core.IdGenesisFromIdenState(core.TypeDefault, state.BigInt())
	if err != nil {
		log.Println("Genesis id failed")
		return "", 0, 0, 0, nil, err
	}

	//TODO: dbg
	fmt.Println("clt:")
	fmt.Println(clt.Root().Hex())
	fmt.Println("state:")
	fmt.Println(state.Hex())
	fmt.Println("did:")
	fmt.Println(hexutil.Encode(did.Bytes()))

	return did.String(), cltId, rotId, torId, state, nil
}

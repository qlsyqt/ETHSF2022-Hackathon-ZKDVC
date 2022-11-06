package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/sql"
	"issuerserver/common"
	"issuerserver/models/database"
	"issuerserver/pkg/db"
	"issuerserver/store"
	"log"
	"math/big"
	"math/rand"
	"time"
)

func GetDidByName(username string) (string, error) {
	didData, err := store.SelectDidByName(username)
	if err != nil {
		return "", err
	}
	return didData.Did, nil
}

func CreateIdentity(username string) (string, error) {
	existed, err := store.SelectDidByName(username)
	if err == nil && existed != nil {
		return "", errors.New("did with duplicate user")
	}
	//1. Create auth claim as the very first claim under the identity
	babyJubjubPrivKey := babyjub.NewRandPrivKey()
	babyJubjubPubKey := babyJubjubPrivKey.Public()
	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	revNonce := crypto.Keccak256([]byte(username))
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(babyJubjubPubKey.X, babyJubjubPubKey.Y),
		core.WithRevocationNonce(new(big.Int).SetBytes(revNonce).Uint64()))
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
		Username:       username,
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
	return did, nil
}

func initAndSaveIdentityTrees(ctx context.Context, authClaim *core.Claim) (string, uint64, uint64, uint64, *merkletree.Hash, error) {
	// Create empty Claims tree
	cltId := rand.Uint64() & 0x7111111111111111 //uint64 with highest bit set to 1 cannot be saved by datbase/sql
	cltStore := sql.NewSqlStorage(db.GetPgDb(), cltId)
	clt, err := merkletree.NewMerkleTree(ctx, cltStore, 32)
	if err != nil {
		log.Println("Claim tree create failed")
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
	if err != nil {
		log.Println("Genesis state failed failed")
		return "", 0, 0, 0, nil, err
	}
	fmt.Println("dbg-createidentity")
	fmt.Println(clt.Root().BigInt(), rot.Root().BigInt(), tor.Root().BigInt())
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

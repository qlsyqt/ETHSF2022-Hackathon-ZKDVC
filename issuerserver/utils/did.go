package utils

import (
	"context"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
	"math/rand"
)

func RandomDid() *core.ID {
	bjjPrivateKey := babyjub.NewRandPrivKey()
	bjjPublicKey := bjjPrivateKey.Public()

	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	revNonce := rand.Uint64()
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(bjjPublicKey.X, bjjPublicKey.Y),
		core.WithRevocationNonce(revNonce))
	ctx := context.Background()
	clts, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	rots, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	tors, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	hi, hv, _ := authClaim.HiHv()
	clts.Add(ctx, hi, hv)
	genesisState, _ := merkletree.HashElems(clts.Root().BigInt(), rots.Root().BigInt(), tors.Root().BigInt())

	id, _ := core.IdGenesisFromIdenState(core.TypeDefault, genesisState.BigInt())
	return id
}

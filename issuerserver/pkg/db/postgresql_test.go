package db

import (
	"context"
	"fmt"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql"
	mtsql "github.com/iden3/go-merkletree-sql/db/sql"
	"math/big"
	"os"
	"testing"
)

// 测试给pq数据库做一个merkle tree，并添加一些叶子节点。
func TestNew(t *testing.T) {
	os.Chdir("../..")
	// 1. BabyJubJub key
	// generate babyJubjub private key randomly
	babyJubjubPrivKey := babyjub.NewRandPrivKey()

	// generate public key from private key
	babyJubjubPubKey := babyJubjubPrivKey.Public()

	// print public key
	fmt.Println(babyJubjubPubKey)

	// 2. Sparse Merkle Tree

	ctx := context.Background()

	// Tree storage
	store := mtsql.NewSqlStorage(GetPgDb(), 666)

	// Generate a new MerkleTree with 32 levels
	mt, err := merkletree.NewMerkleTree(ctx, store, 32)
	if err != nil {
		panic(err)
	}
	// Add a leaf to the tree with index 1 and value 10
	index1 := big.NewInt(1)
	value1 := big.NewInt(10)
	err = mt.Add(ctx, index1, value1)

	// Add another leaf to the tree
	index2 := big.NewInt(2)
	value2 := big.NewInt(15)
	err = mt.Add(ctx, index2, value2)

	// Proof of membership of a leaf with index 1
	proofExist, value, _ := mt.GenerateProof(ctx, index1, mt.Root())

	fmt.Println("Proof of membership:", proofExist.Existence)
	fmt.Println("Value corresponding to the queried index:", value)

	// Proof of non-membership of a leaf with index 4
	proofNotExist, _, _ := mt.GenerateProof(ctx, big.NewInt(4), mt.Root())

	fmt.Println("Proof of membership:", proofNotExist.Existence)

}

// 测试给pq数据库做一个merkle tree，并添加一些叶子节点。
//func TestNew2(t *testing.T) {
//	os.Chdir("../..")
//	p := GetPgDb()
//	nodes := []mtsql.NodeItem{}
//	p.SelectContext(context.Background(), &nodes, "select * from mt_nodes")
//	fmt.Println(nodes)
//}

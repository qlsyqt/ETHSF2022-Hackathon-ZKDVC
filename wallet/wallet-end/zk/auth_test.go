package zk

import (
	"context"
	"fmt"
	"github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-merkletree-sql"
	"github.com/iden3/go-merkletree-sql/db/memory"
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/verifier"
	"io/ioutil"
	"math/big"
	"testing"
	"wallet-end/pkgs/witness"
	"wallet-end/utils"
)

func TestAuth(t *testing.T) {
	fmt.Println("Hello")
	//从私钥入手, 创建身份
	bjjPrivateHex := "0x76806b8801c1ccbc4cfcc18c595b03a4429af84af82e1f13b6c9eb13518d9213"
	var revNonce uint64 = 1234
	bjjPrivate, bjjPublicKey, err := utils.HexPrivateKeyToBjjKeypair(bjjPrivateHex)
	if err != nil {
		panic(err)
	}
	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")
	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(bjjPublicKey.X, bjjPublicKey.Y),
		core.WithRevocationNonce(uint64(revNonce)))
	ctx := context.Background()
	clts, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	rots, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)
	tors, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	hi, hv, _ := authClaim.HiHv()
	clts.Add(ctx, hi, hv)
	genesisState, _ := merkletree.HashElems(clts.Root().BigInt(), rots.Root().BigInt(), tors.Root().BigInt())

	id, _ := core.IdGenesisFromIdenState(core.TypeDefault, genesisState.BigInt())
	authExistsProof, _, _ := clts.GenerateProof(context.Background(), hi, clts.Root())
	authNonRevProof, _, _ := rots.GenerateProof(context.Background(), new(big.Int).SetUint64(revNonce), rots.Root())
	if err != nil {
		panic(err)
	}
	//签名数据，得到input
	challenge, _ := poseidon.HashBytes([]byte("abc"))
	sig := bjjPrivate.SignPoseidon(challenge)

	authInputs := circuits.AuthInputs{
		BaseConfig: circuits.BaseConfig{},
		ID:         id,
		AuthClaim: circuits.Claim{
			TreeState: circuits.TreeState{
				State:          genesisState,
				ClaimsRoot:     clts.Root(),
				RevocationRoot: rots.Root(),
				RootOfRoots:    tors.Root(),
			},
			Claim: authClaim,
			Proof: authExistsProof,
			NonRevProof: &circuits.ClaimNonRevStatus{
				Proof: authNonRevProof,
			},
		},
		Signature: sig,
		Challenge: challenge,
	}
	inputBytes, err := authInputs.InputsMarshal()
	if err != nil {
		panic(err)
	}
	//Create proof
	// 1. compute witness
	wasm, err := ioutil.ReadFile("auth.wasm")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wasm))
	calculator, err := witness.NewCircom2WitnessCalculator(wasm, false)
	if err != nil {
		fmt.Println("calculator create failed")
		panic(err)
	}

	parsedInput, err := witness.ParseInputs(inputBytes)
	if err != nil {
		panic(err)
	}

	wtns, err := calculator.CalculateWTNSBin(parsedInput, false)
	if err != nil {
		fmt.Println("wtns compute failed")
		panic(err)
	}
	zKey, err := ioutil.ReadFile("auth.zkey")
	if err != nil {
		panic(err)
	}
	//wtns, _ := ioutil.ReadFile("zk/witness.wtns")
	proof, err := prover.Groth16Prover(zKey, wtns)
	if err != nil {
		panic(err)
	}
	//
	vKey, _ := ioutil.ReadFile("auth.vkey")
	err = verifier.VerifyGroth16(*proof, vKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Verify success!")
	}
}

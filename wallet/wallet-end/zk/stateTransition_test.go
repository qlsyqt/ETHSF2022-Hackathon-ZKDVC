package zk

import (
	"context"
	"encoding/hex"
	"encoding/json"
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
	"time"
	"wallet-end/pkgs/witness"
	utils2 "wallet-end/utils"
)

func TestStateTransition(t *testing.T) {
	// 1. BabyJubJub key

	// generate babyJubjub private key randomly
	babyJubjubPrivKey, babyJubjubPubKey, _ := utils2.HexPrivateKeyToBjjKeypair("0x76806b8801c1ccbc4cfcc18c595b03a4429af84af82e1f13b6c9eb13518d9213")

	// print public key
	fmt.Println(babyJubjubPubKey)

	// 2. Sparse Merkle Tree

	ctx := context.Background()

	// 3.1. Create a Generic Claim

	// set claim expriation date to 2361-03-22T00:44:48+05:30
	expire := time.Date(2361, 3, 22, 0, 44, 48, 0, time.UTC)

	// set schema
	ageSchema, _ := core.NewSchemaHashFromHex("2e2d1c11ad3e500de68d7ce16a0a559e")

	// define data slots
	birthday := big.NewInt(19960424)
	documentType := big.NewInt(1)

	// set revocation nonce
	revocationNonce := uint64(1909830690)

	// set ID of the claim subject
	subjectId, _ := core.IDFromString("113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2")

	// create claim
	claim, _ := core.NewClaim(ageSchema, core.WithExpirationDate(expire), core.WithRevocationNonce(revocationNonce), core.WithIndexID(subjectId), core.WithIndexDataInts(birthday, documentType))

	// transform claim from bytes array to json
	claimToMarshal, _ := json.Marshal(claim)

	fmt.Println(string(claimToMarshal))

	// 3.2. Create Auth Claim

	authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")

	// Add revocation nonce. Used to invalidate the claim. This may be a random number in the real implementation.
	revNonce := uint64(1)

	authClaim, _ := core.NewClaim(authSchemaHash,
		core.WithIndexDataInts(babyJubjubPubKey.X, babyJubjubPubKey.Y),
		core.WithRevocationNonce(revNonce))

	authClaimToMarshal, _ := json.Marshal(authClaim)

	fmt.Println(string(authClaimToMarshal))

	// 4.1. Generate identity trees

	// Create empty Claims tree
	clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Create empty Revocation tree
	ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Create empty Roots tree
	rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32)

	// Get the Index of the claim and the Value of the authClaim
	hIndex, hValue, _ := authClaim.HiHv()

	// add auth claim to claims tree with value hValue at index hIndex
	clt.Add(ctx, hIndex, hValue)

	// print the roots
	fmt.Println(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())

	// 4.2. Retrieve identity state

	state, _ := merkletree.HashElems(
		clt.Root().BigInt(),
		ret.Root().BigInt(),
		rot.Root().BigInt())

	fmt.Println("Identity State:", state)

	// 4.3. Retrieve Identifier (ID)

	id, _ := core.IdGenesisFromIdenState(core.TypeDefault, state.BigInt())
	fmt.Println(id.BigInt().String())
	// 5. Issuing Claim by Signature
	// 6. Issuing Claim by adding it to the Merkle Tree

	// GENESIS STATE:

	// 1. Generate Merkle Tree Proof for authClaim at Genesis State
	authMTPProof, _, _ := clt.GenerateProof(ctx, hIndex, clt.Root())

	// 2. Generate the Non-Revocation Merkle tree proof for the authClaim at Genesis State
	authNonRevMTPProof, _, _ := ret.GenerateProof(ctx, new(big.Int).SetUint64(revNonce), ret.Root())

	// Snapshot of the Genesis State
	genesisTreeState := circuits.TreeState{
		State:          state,
		ClaimsRoot:     clt.Root(),
		RevocationRoot: ret.Root(),
		RootOfRoots:    rot.Root(),
	}
	// STATE 1:

	// Before updating the claims tree, add the claims tree root at Genesis state to the Roots tree.
	rot.Add(ctx, clt.Root().BigInt(), big.NewInt(0))

	// Create a new random claim
	schemaHex := hex.EncodeToString([]byte("myAge_test_claim"))
	schema, _ := core.NewSchemaHashFromHex(schemaHex)

	code := big.NewInt(51)

	newClaim, _ := core.NewClaim(schema, core.WithIndexDataInts(code, nil))

	// Get hash Index and hash Value of the new claim
	hi, hv, _ := newClaim.HiHv()

	// Add claim to the Claims tree
	clt.Add(ctx, hi, hv)

	// Fetch the new Identity State
	newState, _ := merkletree.HashElems(
		clt.Root().BigInt(),
		ret.Root().BigInt(),
		rot.Root().BigInt())

	fmt.Println(state.BigInt().String(), " :old")
	fmt.Println(newState.BigInt().String(), " :new")

	// Sign a message (hash of the genesis state + the new state) using your private key
	hashOldAndNewStates, _ := poseidon.Hash([]*big.Int{state.BigInt(), newState.BigInt()})

	signature := babyJubjubPrivKey.SignPoseidon(hashOldAndNewStates)

	// Generate state transition inputs
	//通过go-circuits构建输入
	stateTransitionInputs := circuits.StateTransitionInputs{
		ID:                id,
		OldTreeState:      genesisTreeState,
		NewState:          newState,
		IsOldStateGenesis: true,
		AuthClaim: circuits.Claim{
			Claim: authClaim,
			Proof: authMTPProof,
			NonRevProof: &circuits.ClaimNonRevStatus{
				Proof: authNonRevMTPProof,
			},
		},
		Signature: signature,
	}

	// Perform marshalling of the state transition inputs
	inputBytes, _ := stateTransitionInputs.InputsMarshal()
	fmt.Println(string(inputBytes))
	//fmt.Println(inputBytes)
	//通过go-witness-calculator创建见证

	wasm, err := ioutil.ReadFile("stateTransition.wasm")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wasm))
	calculator, err := witness.NewCircom2WitnessCalculator(wasm, false)
	if err != nil {
		fmt.Println("calculator create failed")
		panic(err)
	}

	input, err := witness.ParseInputs(inputBytes)
	if err != nil {
		panic(err)
	}

	wtns, err := calculator.CalculateWTNSBin(input, false)
	if err != nil {
		fmt.Println("wtns compute failed")
		panic(err)
	}
	zKey, err := ioutil.ReadFile("stateTransition.zkey")
	if err != nil {
		panic(err)
	}
	//wtns, _ := ioutil.ReadFile("zk/witness.wtns")
	p, err := prover.Groth16Prover(zKey, wtns)
	if err != nil {
		panic(err)
	}
	//
	vKey, _ := ioutil.ReadFile("stateTransition.vkey")
	err = verifier.VerifyGroth16(*p, vKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("验证成功!")
	}
	//通过go-rapidsnark-prover创建证明
	fmt.Println(p.Proof.A)
	fmt.Println(p.Proof.B)
	fmt.Println(p.Proof.C)
	//通过go-prover-verifier验证证明
}

package utils

import (
	"fmt"
	"github.com/iden3/go-circuits"
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/types"
	"io/ioutil"
	"issuerserver/pkg/zkp/witness"
	"math/big"
)

func GetStateTransitionProof(stateTransitionInputs *circuits.StateTransitionInputs) (*types.ZKProof, error) {
	// Perform marshalling of the state transition inputs
	inputBytes, _ := stateTransitionInputs.InputsMarshal()
	fmt.Println(string(inputBytes))
	//fmt.Println(inputBytes)
	//通过go-witness-calculator创建见证

	wasm, err := ioutil.ReadFile("zk/stateTransition.wasm")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wasm))
	calculator, err := witness.NewCircom2WitnessCalculator(wasm, false)
	if err != nil {
		fmt.Println("calculator create failed")
		return nil, err
	}

	input, err := witness.ParseInputs(inputBytes)
	if err != nil {
		return nil, err
	}

	wtns, err := calculator.CalculateWTNSBin(input, false)
	if err != nil {
		fmt.Println("wtns compute failed")
		return nil, err
	}
	zKey, err := ioutil.ReadFile("zk/stateTransition.zkey")
	if err != nil {
		return nil, err
	}
	//wtns, _ := ioutil.ReadFile("zk/witness.wtns")
	p, err := prover.Groth16Prover(zKey, wtns)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetBigIntABC(proof *types.ZKProof) ([2]*big.Int, [2][2]*big.Int, [2]*big.Int) {
	a := [2]*big.Int{}
	b := [2][2]*big.Int{}
	c := [2]*big.Int{}

	a[0] = StringToBigInt(proof.Proof.A[0])
	a[1] = StringToBigInt(proof.Proof.A[1])

	b[0][0] = StringToBigInt(proof.Proof.B[0][1]) //Order should be reversed(no idea if it a bug of iden3)
	b[0][1] = StringToBigInt(proof.Proof.B[0][0])
	b[1][0] = StringToBigInt(proof.Proof.B[1][1])
	b[1][1] = StringToBigInt(proof.Proof.B[1][0])

	c[0] = StringToBigInt(proof.Proof.C[0])
	c[1] = StringToBigInt(proof.Proof.C[1])

	return a, b, c

}

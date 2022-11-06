package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"wallet-end/pkgs/eths"

	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/iden3/go-rapidsnark/prover"
	zktypes "github.com/iden3/go-rapidsnark/types"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/iden3/go-circuits"
	"io/ioutil"
	"math/big"
	"wallet-end/config"
	stateStub "wallet-end/pkgs/eths/state"
	"wallet-end/pkgs/witness"
	"wallet-end/utils"
)

func SyncIdentityState(identity string, inputs *circuits.StateTransitionInputs) error {
	//Create state transition proof
	cfg := config.GetConfig()
	proof, err := getStateTransitionProof(inputs, cfg)
	a, b, c := getBigIntABC(proof)
	//Call contract "transiteState" function
	ethclient := eths.GetEthClient()

	targetAddr := common.HexToAddress(cfg.Blockchain.StateContract)
	stateInstance, err := stateStub.NewState(targetAddr, ethclient)
	if err != nil {
		return err
	}
	//TxOpts
	ctx := context.Background()
	wallet := common.HexToAddress(utils.AddressFromPrivateKey(cfg.PolygonWallet.MainPrivateKey))
	fmt.Println(wallet)
	nonce, err := ethclient.PendingNonceAt(ctx, wallet)
	if err != nil {
		return err
	}
	fmt.Println("nonce ", nonce)
	nonceBig := new(big.Int)
	nonceBig.SetUint64(nonce)

	chainID, err := ethclient.ChainID(ctx)
	if err != nil {
		return err
	}

	gasPrice, gasLimit, err := utils.EstimateGasPriceAndFee(ethclient, targetAddr)
	if err != nil {
		return err
	}

	txOpts := &bind.TransactOpts{
		From:  wallet,
		Nonce: nonceBig,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			txSigner := types.LatestSignerForChainID(chainID)
			pkey, err := utils.HexPrivateKeyToECDSAKey(cfg.PolygonWallet.MainPrivateKey)
			if err != nil {
				return nil, err
			}
			signedTx, err := types.SignTx(tx, txSigner, pkey)
			if err != nil {
				return nil, err
			}
			return signedTx, nil
		},
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}

	tx, err := stateInstance.TransitState(
		txOpts,
		inputs.ID.BigInt(),
		inputs.OldTreeState.State.BigInt(),
		inputs.NewState.BigInt(),
		inputs.IsOldStateGenesis,
		a,
		b,
		c,
	)
	//TODO: maybe check tx status?
	if err != nil {
		return err
	}

	fmt.Printf("identity state updated. tx hash: %s", tx.Hash().String())
	return nil
}

func getStateTransitionProof(stateTransitionInputs *circuits.StateTransitionInputs, cfg *config.Config) (*zktypes.ZKProof, error) {
	// Perform marshalling of the state transition inputs
	inputBytes, _ := stateTransitionInputs.InputsMarshal()
	fmt.Println(string(inputBytes))

	stCfg := cfg.Zkp["stateTransition"]
	wasm, err := ioutil.ReadFile(stCfg.WasmPath)
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
	zKey, err := ioutil.ReadFile(stCfg.ZkeyPath)
	if err != nil {
		panic(err)
	}
	//wtns, _ := ioutil.ReadFile("zk/witness.wtns")
	p, err := prover.Groth16Prover(zKey, wtns)
	if err != nil {
		panic(err)
	}
	return p, nil
}

func getBigIntABC(proof *zktypes.ZKProof) ([2]*big.Int, [2][2]*big.Int, [2]*big.Int) {
	a := [2]*big.Int{}
	b := [2][2]*big.Int{}
	c := [2]*big.Int{}

	a[0] = utils.StringToBigInt(proof.Proof.A[0])
	a[1] = utils.StringToBigInt(proof.Proof.A[1])

	b[0][0] = utils.StringToBigInt(proof.Proof.B[0][1]) //Order should be reversed(no idea if it a bug of iden3)
	b[0][1] = utils.StringToBigInt(proof.Proof.B[0][0])
	b[1][0] = utils.StringToBigInt(proof.Proof.B[1][1])
	b[1][1] = utils.StringToBigInt(proof.Proof.B[1][0])

	c[0] = utils.StringToBigInt(proof.Proof.C[0])
	c[1] = utils.StringToBigInt(proof.Proof.C[1])

	return a, b, c

}

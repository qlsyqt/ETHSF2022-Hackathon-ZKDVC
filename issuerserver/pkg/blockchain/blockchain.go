package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iden3/go-circuits"
	"issuerserver/config"
	"issuerserver/pkg/eth"
	stateStub "issuerserver/pkg/eth/state"
	"issuerserver/utils"
	"math/big"
)

func SyncIdentityState(identity string, inputs *circuits.StateTransitionInputs) error {
	//Create state transition proof
	proof, err := utils.GetStateTransitionProof(inputs)
	if err != nil {
		return err
	}
	a, b, c := utils.GetBigIntABC(proof)
	//Call contract "transiteState" function
	ethclient := eth.GetEthClient()
	cfg := config.GetConfig()
	targetAddr := common.HexToAddress(cfg.Blockchain.StateContract)
	stateInstance, err := stateStub.NewState(targetAddr, ethclient)
	if err != nil {
		return err
	}
	//TxOpts
	ctx := context.Background()
	wallet := common.HexToAddress(utils.AddressFromPrivateKey(cfg.Blockchain.PrivateKey))
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
			pkey, err := utils.HexPrivateKeyToECDSAKey(cfg.Blockchain.PrivateKey)
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
	fmt.Println("send transaction")
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

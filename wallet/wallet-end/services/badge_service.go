package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	core "github.com/iden3/go-iden3-core"
	"log"
	"math/big"
	"wallet-end/badge"
	"wallet-end/config"
	"wallet-end/models/response"
	"wallet-end/pkgs/codec"
	"wallet-end/pkgs/eths"
	"wallet-end/store"
	"wallet-end/utils"
)

func FetchBadges() ([]*response.FetchBadgeVO, error) {
	//创建合约访问口
	cfg := config.GetConfig()
	client := eths.GetEthClient()

	contractAddress := common.HexToAddress(cfg.Blockchain.BadgeContract)
	badgeContract, err := badge.NewBadgeCaller(contractAddress, client)
	if err != nil {
		return nil, err
	}
	//先调用token.balanceOf获取用户持有徽章数量
	claimerAddress := common.HexToAddress(utils.AddressFromPrivateKey(cfg.PolygonWallet.AuxPrivateKey))
	count, err := badgeContract.BalanceOf(nil, claimerAddress)
	if err != nil {
		return nil, err
	}
	//再调用token.tokenOfOwnerByIndex一个个遍历拿到tokenID即可
	var i int64 = 0
	ans := make([]*response.FetchBadgeVO, 0)
	for i < int64(count.Uint64()) {
		token, err := badgeContract.TokenOfOwnerByIndex(nil, claimerAddress, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		dataCategory, err := badgeContract.GetTokenCategory(nil, token)
		if err != nil {
			return nil, err
		}
		ans = append(ans, &response.FetchBadgeVO{
			DataCategory:  dataCategory,
			Badge:         token.Text(10),
			BadgeContract: cfg.Blockchain.BadgeContract,
		})
		i += 1
	}
	return ans, nil
}

func ClaimBadge(hIndex string) error {
	//获取claim
	hIndexBytes, err := hexutil.Decode(hIndex)
	if err != nil {
		return err
	}
	dbClaim, err := store.SelectClaimByHIndex(new(big.Int).SetBytes(hIndexBytes))
	if err != nil {
		return err
	}
	claim := &core.Claim{}
	err = claim.UnmarshalBinary(dbClaim.ClaimBinary)
	if err != nil {
		return err
	}
	//解析出dcp
	dcp := codec.DecodeDcpClaim(claim)
	//创建合约访问口
	cfg := config.GetConfig()
	ethclient := eths.GetEthClient()
	contractAddress := common.HexToAddress(cfg.Blockchain.BadgeContract)
	badgeContract, err := badge.NewBadge(contractAddress, ethclient)
	if err != nil {
		return err
	}
	//先调用token.balanceOf获取用户持有徽章数量
	claimerAddress := common.HexToAddress(utils.AddressFromPrivateKey(cfg.PolygonWallet.AuxPrivateKey))

	//txOpts
	nonce, err := ethclient.PendingNonceAt(context.Background(), claimerAddress)
	if err != nil {
		return err
	}
	fmt.Println("nonce ", nonce)
	nonceBig := new(big.Int)
	nonceBig.SetUint64(nonce)

	chainID, err := ethclient.ChainID(context.Background())
	if err != nil {
		return err
	}

	gasPrice, gasLimit, err := utils.EstimateGasPriceAndFee(ethclient, contractAddress)
	if err != nil {
		return err
	}

	txOpts := &bind.TransactOpts{
		From:  claimerAddress,
		Nonce: nonceBig,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			txSigner := types.LatestSignerForChainID(chainID)
			pkey, err := utils.HexPrivateKeyToECDSAKey(cfg.PolygonWallet.AuxPrivateKey)
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

	tx, err := badgeContract.MintNative(txOpts, dcp.DataCategory, dcp.SubCategory)
	if err != nil {
		return err
	}
	log.Printf("tx hash: %s", tx.Hash().Hex())
	return nil
}

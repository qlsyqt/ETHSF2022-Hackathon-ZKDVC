package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"issuerserver/pkg/logging"
	"math/big"
	"strings"
)
import "golang.org/x/crypto/sha3"

func HexPrivateKeyToECDSAKey(key string) (*ecdsa.PrivateKey, error) {
	if strings.ToLower(key[:2]) == "0x" {
		key = key[2:]
	}
	pk, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func AddressFromPrivateKey(privateKeyHex string) string {
	if strings.ToLower(privateKeyHex[:2]) == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logging.Error.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	var buf []byte

	hash := sha3.NewLegacyKeccak256() //do not use New256() on eth and polygon
	hash.Write(publicKeyBytes[1:])    // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address)).Hex()
}

func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256() //do not use New256() on eth and polygon
	hash.Write(publicKey[1:])         // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

func RandomECDSAKeyPair() (string, common.Address) {
	privateKey, _ := crypto.GenerateKey()
	pKeyBytes := crypto.FromECDSA(privateKey)
	pKeyStr := hexutil.Encode(pKeyBytes)

	pubKey := privateKey.PublicKey
	addr := crypto.PubkeyToAddress(pubKey)

	return pKeyStr, addr
}

func EstimateGasPriceAndFee(ethclient *ethclient.Client, targetAddr common.Address) (*big.Int, uint64, error) {
	//bn, err := ethclient.BlockNumber(context.Background())
	//if err != nil {
	//	return nil, 0, err
	//}
	//bignumBn := big.NewInt(0).SetUint64(bn)
	//blk, err := ethclient.BlockByNumber(context.Background(), bignumBn)
	//if err != nil {
	//	return nil, 0, err
	//}
	//config := params.MainnetChainConfig //TODO
	//baseFee := misc.CalcBaseFee(config, blk.Header())

	//estimatedGas, err := ethclient.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &targetAddr,
	//	Data: []byte{0},
	//})
	//if err != nil {
	//	return nil, 0, err
	//}
	//TODO : 写入配置，或者研究下标准做法例如hardhat
	baseFee := new(big.Int).SetUint64(1294839261 * 2)
	estimatedGas := uint64(1000000)
	return baseFee, estimatedGas, nil
}

func HashChallengeAndWallet(challenge []byte, walletAddress common.Address) (*big.Int, error) {
	preimage := make([]byte, len(challenge)+20)

	copy(preimage, challenge)
	copy(preimage[len(challenge):], walletAddress.Bytes())

	digestBigInt, err := poseidon.HashBytes(preimage)
	if err != nil {
		return nil, err
	}
	return digestBigInt, nil
}

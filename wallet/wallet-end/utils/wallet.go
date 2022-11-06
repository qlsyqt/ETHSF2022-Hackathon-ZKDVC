package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
)

func AddressFromPrivateKey(privateKeyHex string) string {
	if strings.ToLower(privateKeyHex[:2]) == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	var buf []byte

	hash := sha3.NewLegacyKeccak256() //do not use New256() on eth and polygon
	hash.Write(publicKeyBytes[1:])    // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address)).Hex()
}

func PublicKeyFromPrivateKey(privateKeyHex string) *ecdsa.PublicKey {
	if strings.ToLower(privateKeyHex[:2]) == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return publicKeyECDSA
}

func SignDigest(msgHash []byte, privateKeyHex string) []byte {
	if strings.ToLower(privateKeyHex[:2]) == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)

	sig, _ := crypto.Sign(msgHash, privateKey)
	//65 bytes
	return sig
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

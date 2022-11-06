package utils

import (
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/utils"
)

func HexPrivateKeyToBjjKeypair(key string) (babyjub.PrivateKey, *babyjub.PublicKey, error) {
	privateKeyBytes, err := utils.HexDecode(key)
	if err != nil {
		return [32]byte{}, nil, err
	}
	var privateKey babyjub.PrivateKey = [32]byte{}
	copy(privateKey[:], privateKeyBytes)
	publicKey := privateKey.Public()

	return privateKey, publicKey, nil
}

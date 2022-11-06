package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPublicKeyBytesToAddress(t *testing.T) {
	//Generate signature
	privateKey, _ := crypto.GenerateKey()
	message := []byte("hello")
	hash := crypto.Keccak256Hash(message)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	addr := PublicKeyBytesToAddress(publicKeyBytes)
	//Verify

	signature, _ := crypto.Sign(hash.Bytes(), privateKey)

	//Verify signature
	recoverBytes, _ := crypto.Ecrecover(hash.Bytes(), signature)
	recoverAddr := PublicKeyBytesToAddress(recoverBytes)

	assert.True(t, addr == recoverAddr)

}

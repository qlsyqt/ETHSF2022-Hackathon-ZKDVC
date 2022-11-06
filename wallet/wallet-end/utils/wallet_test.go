package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAddressFromPrivateKey(t *testing.T) {
	key, _ := crypto.GenerateKey()
	bytes := crypto.FromECDSA(key)
	fmt.Println(hexutil.Encode(bytes))
}

func TestSign(t *testing.T) {
	msg := []byte("zk-airdrop")
	hash := crypto.Keccak256Hash(msg)

	sig := SignDigest(hash.Bytes(), "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a") //need to cut the sig[64]
	assert.Equal(t, 64, len(sig))
}

func TestPublicKeyFromPrivateKeyWith64(t *testing.T) {
	public := PublicKeyFromPrivateKey("5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a") //need to cut the sig[64]

	arr1 := BigIntToTuple(public.X)
	arr2 := BigIntToTuple(public.Y)

	fmt.Println(arr1)
	fmt.Println(arr2)

	//assert.Equal(t, 64, len(sig))
}

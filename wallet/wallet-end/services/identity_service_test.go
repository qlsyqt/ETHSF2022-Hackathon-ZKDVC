package services

import (
	"fmt"
	"os"
	"testing"
	"wallet-end/config"
)

func TestFetchIdentity(t *testing.T) {
	os.Chdir("..")
	cfg := config.GetConfig()
	expected, _ := deduceIdentity(cfg.PolygonId.PrivateKey, cfg.PolygonId.AuthRevocationNonce)

	did, err := FetchIdentity()
	if err != nil {
		panic(err)
	}
	fmt.Println(expected.String())
	fmt.Println(did)
	fmt.Println(myDid)
}

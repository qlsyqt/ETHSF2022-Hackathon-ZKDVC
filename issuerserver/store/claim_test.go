package store

import (
	"fmt"
	"issuerserver/models/database"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestInsertAndSelectClaim(t *testing.T) {
	os.Chdir("..")
	fmt.Println(os.Getwd())
	b := big.NewInt(1111).Bytes()
	c := database.ClaimData{
		ClaimHi:     b,
		ClaimBinary: []byte("aaabbb"),
		Issuer:      "issuer",
		Revoked:     false,
		ClaimType:   "dcp",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := InsertClaim(&c)
	if err != nil {
		panic(err)
	}

	c.Id = 0
	c.ClaimHi = big.NewInt(2222).Bytes()
	err = InsertClaim(&c)
	if err != nil {
		panic(err)
	}

	c.Id = 0
	c.ClaimHi = big.NewInt(3333).Bytes()
	err = InsertClaim(&c)
	if err != nil {
		panic(err)
	}

	c.Id = 0
	c.Revoked = true
	c.ClaimHi = big.NewInt(4444).Bytes()
	err = InsertClaim(&c)
	if err != nil {
		panic(err)
	}

	c.Id = 0
	c.Revoked = false
	c.ClaimType = "auth"
	c.ClaimHi = big.NewInt(5555).Bytes()
	err = InsertClaim(&c)
	if err != nil {
		panic(err)
	}
	r, err := SelectClaimByHIndex(big.NewInt(1111))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	batch, err := SelectClaimsByBatch(-1, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(batch[0].Id, batch[1].Id)
}

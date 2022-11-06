package codec

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/go-playground/assert/v2"
	core "github.com/iden3/go-iden3-core"
	"issuerserver/models/common"
	"testing"
)

func TestDcpEncodeAndDecode(t *testing.T) {
	dcp := &common.DCP{
		DataCategory: "nft",
		SubCategory:  "0xc942875c9Ed96E9dBa8eFb9F88C99EfcD3FED908",
		LowerBoundary: common.Boundary{
			Include: true,
			Value:   7,
		},
		UpperBoundary: common.Boundary{
			Include: false,
			Value:   7,
		},
		HolderAddress: "0xc942875c9Ed96E9dBa8eFb9F88C99EfcD3FED908",
	}
	claim := core.Claim{}
	err := EncodeDcpClaim(&claim, dcp)
	if err != nil {
		fmt.Println(err)
	}

	dcpr := DecodeDcpClaim(&claim)

	fmt.Println(dcpr)

	dcp.DataCategory = "snapshot"
	dcp.SubCategory = "cyz.eth"

	err = EncodeDcpClaim(&claim, dcp)
	if err != nil {
		panic(err)
	}

	dcpr = DecodeDcpClaim(&claim)

	fmt.Println(dcpr)
}

func TestStringEncodeDecode(t *testing.T) {
	str := "NFT"
	slot := encodeUtf8String(str)
	fmt.Println(slot)

	recovered := decodeUtf8String(slot)
	assert.Equal(t, str, recovered)
}

func TestAddressEncodeDecode(t *testing.T) {
	addr := "0xc942875c9Ed96E9dBa8eFb9F88C99EfcD3FED908"
	slot := encodeEthAddress(addr)
	fmt.Println(slot)

	recovered := decodeEthAddress(slot)
	fmt.Println(recovered)
	assert.Equal(t, addr, recovered)
}

func TestBoundariesEncodeDecode2(t *testing.T) {
	lb := &common.Boundary{
		Include: false,
		Value:   math.MinInt32,
	}
	ub := &common.Boundary{
		Include: false,
		Value:   math.MaxInt32,
	}
	slot := encodeBoundaries(lb, ub)
	fmt.Println(slot)

	lbr, ubr := decodeBoundaries(slot)
	assert.Equal(t, lb.Include, lbr.Include)
	assert.Equal(t, lb.Value, lbr.Value)
	assert.Equal(t, ubr.Include, ubr.Include)
	assert.Equal(t, ubr.Value, ubr.Value)
}

func TestBoundariesEncodeDecode(t *testing.T) {
	lb := &common.Boundary{
		Include: true,
		Value:   7,
	}
	ub := &common.Boundary{
		Include: false,
		Value:   10,
	}
	slot := encodeBoundaries(lb, ub)
	fmt.Println(slot)

	lbr, ubr := decodeBoundaries(slot)
	assert.Equal(t, lb.Include, lbr.Include)
	assert.Equal(t, lb.Value, lbr.Value)
	assert.Equal(t, ubr.Include, ubr.Include)
	assert.Equal(t, ubr.Value, ubr.Value)
}

//
//func TestBoundries(t *testing.T) {
//	b := encodeBoundry(true, 30)
//	expected := [32]byte{0, 0, 0, 30, 1}
//	fmt.Println(b)
//	assert.True(t, bytes.Compare(expected[:], b[:]) == 0)
//	include, value := decodeBoundry(b)
//	assert.Equal(t, include, true)
//	assert.Equal(t, value, int32(30))
//}

func TestBinary(t *testing.T) {
	b := 0b00000011
	fmt.Println(b)
}

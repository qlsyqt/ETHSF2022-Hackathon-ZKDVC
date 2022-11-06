package dvc

import (
	"github.com/ethereum/go-ethereum/common"
)

func nftEncodeSubcategory(subCategory string) []byte {
	return common.FromHex(subCategory)
}

func nftDecodeSubcategory(subCategory []byte) string {
	return common.BytesToAddress(subCategory).String()
}

package dvc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"issuerserver/models/dvc"
	"strconv"
	"strings"
)

func nftQuery(ctx context.Context, subCategory string, wallet string) (int, error) {

	//Send request
	bodyBytes, err := sendRequestAndGetResponse(ctx, "nft", subCategory, wallet)

	fmt.Println("dvc returns ", string(bodyBytes))
	if err != nil {
		if err.Error() == "404 not found" {
			return 0, nil
		}
		return 0, err
	}
	//Find matching subcategory item
	walletData := make([]dvc.NftItem, 0)
	err = json.Unmarshal(bodyBytes, &walletData)
	if err != nil {
		return 0, err
	}
	if strings.ToLower(subCategory[:2]) == "0x" {
		subCategory = subCategory[2:]
	}
	subCategory = strings.ToLower(subCategory)
	for _, item := range walletData {
		if strings.ToLower(item.Contract[:2]) == "0x" {
			item.Contract = item.Contract[2:]
		}
		fmt.Printf("%s vs %s\n", item.Contract, subCategory)
		if strings.ToLower(item.Contract) == subCategory {
			ans, err := strconv.Atoi(item.Count)
			if err != nil {
				return 0, err
			}
			return ans, nil
		}
	}
	return 0, nil
}

func nftEncodeSubcategory(subCategory string) []byte {
	return common.FromHex(subCategory)
}

func nftDecodeSubcategory(subCategory []byte) string {
	return common.BytesToAddress(subCategory).String()
}

package dvc

import (
	"context"
	"encoding/json"
	"issuerserver/models/dvc"
)

func ensQuery(ctx context.Context, subCategory string, wallet string) (int, error) {
	//Send request
	bodyBytes, err := sendRequestAndGetResponse(ctx, "ens", subCategory, wallet)
	if err != nil {
		if err.Error() == "404 not found" {
			return 0, nil
		}
		return 0, err
	}
	//Find matching subcategory item
	walletData := dvc.EnsItem{}

	err = json.Unmarshal(bodyBytes, &walletData)
	if err != nil {
		return 0, err
	}
	return walletData.HoldCount, nil
}

// ENS does not have subcategory
func ensEncodeSubcategory(subCategory string) []byte {
	return []byte{}
}

// ENS does not have subcategory
func ensDecodeSubcategory(subCategory []byte) string {
	return ""
}

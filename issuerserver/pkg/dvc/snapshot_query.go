package dvc

import (
	"context"
	"encoding/json"
	"errors"
)

func snapshotQuery(ctx context.Context, subCategory string, wallet string) (int, error) {
	//Send request
	bodyBytes, err := sendRequestAndGetResponse(ctx, "snapshot", subCategory, wallet)
	if err != nil {
		if err.Error() == "404 not found" {
			return 0, nil
		}
		return 0, err
	}
	//Find matching subcategory item
	walletData := make(map[string]int, 0)
	err = json.Unmarshal(bodyBytes, &walletData)
	if err != nil {
		return 0, errors.New(string(bodyBytes))
	}
	if len(subCategory) == 0 {
		return walletData["total"], nil
	}
	if count, ok := walletData[subCategory]; ok {
		return count, nil
	}
	return 0, nil
}

func snapshotEncodeSubcategory(subCategory string) []byte {
	return []byte(subCategory)
}

func snapshotDecodeSubcategory(subCategory []byte) string {
	return string(subCategory)
}

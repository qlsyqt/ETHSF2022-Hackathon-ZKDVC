package dvc

//TODO：pagerank的返回格式需要请教下各位
//func pageRankQuery(ctx context.Context, subCategory string, wallet string) (int, error) {
//	//Send request
//	bodyBytes, err := sendRequestAndGetResponse(ctx, "pagerank", subCategory, wallet)
//	//Find matching subcategory item
//	walletData := dvc.EnsItem{}
//	err = json.Unmarshal(bodyBytes, &walletData)
//	if err != nil {
//		return 0, err
//	}
//	return walletData.HoldCount, nil
//}
//
//// Pagerank does not have subcategory
//func pageRankEncodeSubcategory(subCategory string) []byte {
//	return []byte{}
//}
//
//// Pagerank does not have subcategory
//func pageRankDecodeSubcategory(subCategory []byte) string {
//	return ""
//}

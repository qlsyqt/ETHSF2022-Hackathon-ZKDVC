package dvc

func snapshotEncodeSubcategory(subCategory string) []byte {
	return []byte(subCategory)
}

func snapshotDecodeSubcategory(subCategory []byte) string {
	return string(subCategory)
}

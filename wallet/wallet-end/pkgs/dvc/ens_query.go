package dvc

// ENS does not have subcategory
func ensEncodeSubcategory(subCategory string) []byte {
	return []byte{}
}

// ENS does not have subcategory
func ensDecodeSubcategory(subCategory []byte) string {
	return ""
}

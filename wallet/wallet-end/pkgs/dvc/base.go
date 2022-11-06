package dvc

//type DvcQuery interface {
//	Query(ctx context.Context, subCategory string, wallet string) (int, error)
//}

var encodeStubs map[string]func(subCategory string) []byte
var decodeStubs map[string]func(subCategory []byte) string

func EncodeSubcategory(dataCategory, subCategory string) []byte {
	if f, ok := encodeStubs[dataCategory]; ok {
		return f(subCategory)
	}
	return []byte{}
}

func DecodeSubcategory(dataCategory string, subCategory []byte) string {
	if f, ok := decodeStubs[dataCategory]; ok {
		return f(subCategory)
	}
	return ""
}

func init() {

	encodeStubs = make(map[string]func(subCategory string) []byte)
	encodeStubs["0"] = ensEncodeSubcategory
	encodeStubs["1"] = nftEncodeSubcategory
	encodeStubs["2"] = snapshotEncodeSubcategory

	decodeStubs = make(map[string]func(subCategory []byte) string)
	decodeStubs["0"] = ensDecodeSubcategory
	decodeStubs["1"] = nftDecodeSubcategory
	decodeStubs["2"] = snapshotDecodeSubcategory

}

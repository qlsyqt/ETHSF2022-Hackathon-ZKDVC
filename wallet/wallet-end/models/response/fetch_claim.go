package response

type FetchClaimVO struct {
	HIndex       string `json:"hIndex"`
	DataCategory string `json:"dataCategory"`
	SubCategory  string `json:"subCategory"`
	Interval     string `json:"interval"`
	//             hIndex: '1',
	//             dataCategory: 'nft',
	//             subCategory: '0xaaaabbbccc',
	//             interval: "[0, 30)"
}

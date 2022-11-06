package utils

import (
	"github.com/ethereum/go-ethereum/common/math"
	"issuerserver/models/common"
)

func InRange(intValue int, lowerBoundary, upperboundary *common.Boundary) bool {
	value := int32(intValue)
	var lowerSatiesfied bool
	if lowerBoundary.Include {
		lowerSatiesfied = value >= lowerBoundary.Value
	} else {
		lowerSatiesfied = value > lowerBoundary.Value
	}

	var upperSatiesfied bool
	if upperboundary.Include {
		upperSatiesfied = value <= upperboundary.Value
	} else {
		upperSatiesfied = value < upperboundary.Value
	}

	return lowerSatiesfied && upperSatiesfied
}

// TODO: 和前端确认。还有都改成整数
// [1, 1, "78", 0]
// 是否选择 ｜如果选择，是否包含 ｜ 界值
func RawToBoundary(bound []int32, lower bool) (*common.Boundary, error) {

	//If select bound
	boundVal := int32(0)
	if bound[0] == 1 {
		boundVal = bound[2]
	} else if lower {
		boundVal = 0
	} else {
		boundVal = math.MaxInt32
	}

	include := false
	//if select bound, respect to bound[1]
	if bound[0] == 1 {
		include = bound[1] != 0
	} else {
		include = false
	}

	boundary := common.Boundary{
		Include: include,
		Value:   boundVal,
	}

	return &boundary, nil
}

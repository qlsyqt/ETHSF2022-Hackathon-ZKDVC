package codec

import (
	"fmt"
	"math"
	"testing"
	"wallet-end/models/common"
)

func TestRangeToString(t *testing.T) {

	lb := common.Boundary{
		Include: true,
		Value:   math.MinInt32,
	}

	ub := common.Boundary{
		Include: false,
		Value:   math.MaxInt32,
	}

	fmt.Println(RangeToString(&lb, &ub))
}

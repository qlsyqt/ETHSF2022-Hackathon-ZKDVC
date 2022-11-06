package utils

import (
	"github.com/stretchr/testify/assert"
	"issuerserver/models/common"
	"testing"
)

func TestInRange(t *testing.T) {
	lb := &common.Boundary{
		Include: true,
		Value:   1,
	}

	up := &common.Boundary{
		Include: false,
		Value:   10,
	}

	assert.True(t, InRange(1, lb, up))
	assert.True(t, InRange(2, lb, up))
	assert.True(t, InRange(9, lb, up))
	assert.False(t, InRange(10, lb, up))
	assert.False(t, InRange(11, lb, up))
}

func TestRawToBoundary(t *testing.T) {
	lb, err := RawToBoundary([]int32{1, 0, 1, 0}, true)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, false, lb.Include)
	assert.Equal(t, int32(1), lb.Value)
	ub, err := RawToBoundary([]int32{1, 1, 10, 0}, false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, true, ub.Include)
	assert.Equal(t, int32(10), ub.Value)
	//[0, 0, 0, 1]
	b, err := RawToBoundary([]int32{0, 1, 10, 1}, true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, false, b.Include)
	assert.Equal(t, int32(0), b.Value)
}

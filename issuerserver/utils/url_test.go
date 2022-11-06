package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	s, err := Concat("https://www.aaa.com", "bbb", "ccc")
	assert.True(t, err == nil)
	assert.Equal(t, s, "https://www.aaa.com/bbb/ccc")
}

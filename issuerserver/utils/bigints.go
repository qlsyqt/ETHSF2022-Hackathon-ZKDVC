package utils

import (
	"math/big"
)

func Uint8ArrayToBigInt(x []byte) *big.Int {
	ans := big.NewInt(0)

	for _, b := range x {
		//left shit by 8 bits (one byte)
		ans = big.NewInt(0).Mul(ans, big.NewInt(256))
		ans = big.NewInt(0).Add(ans, new(big.Int).SetUint64(uint64(b)))
	}
	return ans
}

// Used to split signature. Split x into k parts, each with at most n bytes
func BigIntToArray(n, k int, x *big.Int) []*big.Int {
	mod := big.NewInt(1)

	for idx := 0; idx <= n; idx++ {
		mod = big.NewInt(0).Mul(mod, big.NewInt(2))
	}

	ans := make([]*big.Int, 0)
	xTemp := x
	for idx := 0; idx < k; idx++ {
		m := big.NewInt(0).Mod(xTemp, mod)
		ans = append(ans, m)
		xTemp = big.NewInt(0).Div(xTemp, mod)
	}

	return ans
}

// Used to split public keys
func BigIntToTuple(x *big.Int) []*big.Int {
	mod, _ := new(big.Int).SetString("77371252455336267181195264", 10)
	ret := make([]*big.Int, 3)

	xTemp := x
	for idx := 0; idx < 3; idx++ {
		ret[idx] = new(big.Int).Mod(xTemp, mod)
		xTemp = new(big.Int).Div(xTemp, mod)
	}
	return ret
}

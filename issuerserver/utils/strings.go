package utils

import "math/big"

func StringToBigInt(str string) *big.Int {

	ans := new(big.Int)
	ans.SetString(str, 10)
	return ans
}

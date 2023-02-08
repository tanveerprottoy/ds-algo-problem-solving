package hackerrank

import "math/big"

func ExtraLongFactorials(n int32) big.Int {
	var res big.Int
	for n > 0 {
		n, res = longFactorials(n, res)
	}
	return res
}

func longFactorials(n int32, res big.Int) (int32, big.Int) {
	if n <= 0 {
		return n, res
	}
	res = *res.Mul(&res, big.NewInt(int64(n)))
	return longFactorials(n-1, res)
}

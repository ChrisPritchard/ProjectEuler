package main

import "math/big"

// convenience methods for big.Int

func bigmul(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(a, b)
	return r
}

func bigmuli(a *big.Int, b int) *big.Int {
	return bigmul(a, big.NewInt(int64(b)))
}

func bigadd(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Add(a, b)
	return r
}

func bigaddi(a *big.Int, b int) *big.Int {
	return bigadd(a, big.NewInt(int64(b)))
}

func bigdiv(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Div(a, b)
	return r
}

func bigdivi(a *big.Int, b int) *big.Int {
	return bigdiv(a, big.NewInt(int64(b)))
}

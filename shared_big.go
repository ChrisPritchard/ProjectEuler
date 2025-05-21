package main

import "math/big"

// convenience methods for big.Int. each *returns* a new big int

// creates a big int from an integer (just avoiding the clunky int64 conversion)
func newbig(n int) *big.Int {
	return big.NewInt(int64(n))
}

// creates a bigint from the string representation of a number (base 10)
func newbig_fromstr(s string) *big.Int {
	r := newbig(0)
	r.SetString(s, 10)
	return r
}

// multiplies two bigints and returns the result
func bigmul(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(a, b)
	return r
}

// multiplies a bigint by an integer and returns the result
func bigmuli(a *big.Int, b int) *big.Int {
	return bigmul(a, newbig(b))
}

// adds two bigints and returns the result
func bigadd(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Add(a, b)
	return r
}

// adds an integer to a bigint and returns the result
func bigaddi(a *big.Int, b int) *big.Int {
	return bigadd(a, big.NewInt(int64(b)))
}

// divides bigint a by bigint b and returns the result (the quotient)
func bigdiv(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Div(a, b)
	return r
}

// divides a bigint a by an integer b and returns the result (the quotient)
func bigdivi(a *big.Int, b int) *big.Int {
	return bigdiv(a, big.NewInt(int64(b)))
}

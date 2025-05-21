package main

import (
	"slices"
)

// calculate all primes in the range 1 to max using the prime sieve method
func prime_sieve(max int) []int {
	res := []int{}
	sieve := make([]bool, max)

	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}

		res = append(res, i)
		for j := i + i; j < len(sieve); j += i {
			sieve[j] = true
		}
	}

	return res
}

// simple integer power function
func pown(base int, power uint) int {
	if power == 0 {
		return 1
	}
	n := base
	for power > 1 {
		n *= base
		power--
	}
	return n
}

// simplem integer abs function
func absi(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// separates a number into its digits, e.g. 123 becomes 1,2,3
func value_to_digits(value int) []int {
	res := []int{}

	for value != 0 {
		res = append(res, value%10)
		value /= 10
	}

	return reverse(res)
}

// converts digits into a number, e.g. 1,2,3 becomes 123
func digits_to_value(digits []int) int {
	sum := 0
	for i, v := range reverse(digits) {
		sum += pown(10, uint(i)) * v
	}
	return sum
}

// returns all possible combinations of the input array
func permute(values []int) [][]int {

	var expander func([]int, []int) [][]int
	expander = func(current []int, rem []int) [][]int {
		results := [][]int{}
		for i, v := range rem {
			next := make([]int, len(current)+1)
			copy(next, current)
			next[len(next)-1] = v
			if len(rem) > 1 {
				new_rem := make([]int, len(rem))
				copy(new_rem, rem)
				new_rem = slices.Delete(new_rem, i, i+1)
				results = append(results, expander(next, new_rem)...)
			} else {
				results = append(results, next)
			}
		}
		return results
	}

	return expander([]int{}, values)
}

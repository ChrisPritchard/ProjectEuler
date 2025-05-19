package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "all" {
		Problems_001_010()
		Problems_011_020()
		Problems_021_030()
		Problems_031_040()
		Problems_041_050()
	}

	problem_051()
}

func problem_051() {
	// get a million primes
	// 8 primes could start at 0, 1 or 2 (3,4,5,6,7,8,9)
	// so go through primes, discarding any without the first three
	// when the first three are found, calculate positions and calculate candidates

	primes := prime_sieve(1_000_000)
	prime_set := NewSet(primes...)

	indexes := func(digits []int, looked_for int) []int {
		res := []int{}
		for i := range digits {
			if digits[i] == looked_for {
				res = append(res, i)
			}
		}
		return res
	}

	setter := func(digits []int, mask []int, replacement int) int {
		for i := range digits {
			if slices.Contains(mask, i) {
				digits[i] = replacement
			}
		}

		return digits_to_value(digits)
	}

	test_mask := func(digits []int, mask []int, start int) bool {
		count := 1
		for j := start; j <= 9; j++ {
			candidate := setter(digits, mask, j)
			if prime_set.Contains(candidate) {
				count++
			}
		}
		return count >= 8
	}

	for i := range primes {
		prime_digits := value_to_digits(primes[i])
		pos_0 := indexes(prime_digits, 0)
		if len(pos_0) > 0 {
			if test_mask(prime_digits, pos_0, 1) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
		pos_1 := indexes(prime_digits, 1)
		if len(pos_1) > 0 {
			if test_mask(prime_digits, pos_1, 2) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
		pos_2 := indexes(prime_digits, 2)
		if len(pos_2) > 0 {
			if test_mask(prime_digits, pos_2, 3) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
	}

}

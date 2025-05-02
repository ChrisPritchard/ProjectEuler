package main

import (
	"os"
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
	extract_key := func(digits []int, positions []int) int {
		transformer := func(i int) int {
			if i < len(digits) {
				return digits[i]
			} else {
				return 0
			}
		}
		just_positions := transform(positions, transformer)
		return digits_to_value(just_positions)
	}

	primes := prime_sieve(1_000_000)
}

package main

import "fmt"

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()
	Problems_031_040()

	problem_041()
}

func problem_041() {
	// divisibility rule by 3 says 8, 9 digit pandigitals are all composite
	primes := prime_sieve(7_654_321)
	for p := len(primes) - 1; p > 0; p-- {
		digits := value_to_digits(primes[p])
		unique := NewSet(digits...)
		if unique.Size() < len(digits) {
			continue
		}

		valid := true
		for i := 1; i <= len(digits); i++ {
			if !unique.Contains(i) {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println("problem 041:", primes[p])
			return
		}
	}
}

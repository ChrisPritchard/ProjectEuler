package main

import "fmt"

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()
	Problems_031_040()

	problem_041()
	problem_042()
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

func problem_042() {
	words, _ := read_words("0042_words.txt")

	is_triangle_number := func(n int) bool {
		i := 1
		for {
			t := (i * (i + 1)) / 2
			if t == n {
				return true
			} else if t > n {
				break
			}
			i++
		}
		return false
	}

	count := 0
	for _, word := range words {
		sum := 0
		for _, c := range word {
			sum += (int(c) - int('A')) + 1
		}

		if is_triangle_number(sum) {
			count++
		}
	}

	fmt.Println("problem 042:", count)
}

package main

import "fmt"

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()
	Problems_031_040()

	problem_041()
	problem_042()
	problem_043()
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

func problem_043() {
	options := permute([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	sum := 0

	check := func(o []int, i int, p int) bool {
		return digits_to_value([]int{o[i-1], o[i], o[i+1]})%p == 0
	}

	for _, o := range options {
		if check(o, 2, 2) && check(o, 3, 3) && check(o, 4, 5) && check(o, 5, 7) && check(o, 6, 11) && check(o, 7, 13) && check(o, 8, 17) {
			sum += digits_to_value(o)
		}
	}

	fmt.Println("problem 043:", sum)
}

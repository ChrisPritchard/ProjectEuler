package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()

	problem_031()
	problem_032()
	problem_033()
	problem_034()
	problem_035()
}

func problem_031() {

	dp := make([]int, 200+1)
	dp[0] = 1
	for _, c := range []int{1, 2, 5, 10, 20, 50, 100, 200} {
		for i := c; i < len(dp); i++ {
			dp[i] += dp[i-c]
		}
	}

	fmt.Println("problem 031:", dp[200])
}

func problem_032() {

	var expander func(n int, current []int) [][]int
	expander = func(n int, current []int) [][]int {
		results := [][]int{}
		for i := range 9 {
			if current[i] != 0 {
				continue
			}
			new_option := make([]int, 9)
			copy(new_option, current)
			new_option[i] = n
			if n == 1 {
				results = append(results, new_option)
			} else {
				results = append(results, expander(n-1, new_option)...)
			}
		}
		return results
	}

	all_permuts := expander(9, make([]int, 9))
	seen := make(map[int]struct{})

	num := func(vals []int) int {
		res := 0
		slices.Reverse(vals)
		for i, v := range vals {
			res += v * int(math.Pow(10., float64(i)))
		}
		return res
	}

	tester := func(set []int) int {

		c := num(set[5 : 8+1])
		if _, exists := seen[c]; exists {
			return 0
		}

		a := set[0]
		b := num(set[1 : 4+1])

		if a*b == c {
			seen[c] = struct{}{}
			return c
		}

		a = num(set[0 : 3+1])
		b = set[4]
		if a*b == c {
			seen[c] = struct{}{}
			return c
		}

		a = num(set[0 : 1+1])
		b = num(set[2 : 4+1])
		if a*b == c {
			seen[c] = struct{}{}
			return c
		}

		a = num(set[0 : 2+1])
		b = num(set[3 : 4+1])
		if a*b == c {
			seen[c] = struct{}{}
			return c
		}

		return 0
	}

	full_sum := 0
	for v := range all_permuts {
		full_sum += tester(all_permuts[v])
	}

	fmt.Println("problem 032:", full_sum)
}

func problem_033() {
	numerators := 1
	denominators := 1

	for i := 1; i < 10; i++ {
		for j := range 10 {
			for k := 1; k < 10; k++ {
				for m := 1; m < 10; m++ {
					if i != j && k != m && j == k && float64(i*10+j)/float64(k*10+m) == (float64(i)/float64(m)) {
						numerators *= i*10 + j
						denominators *= k*10 + m
					}
				}
			}
		}
	}

	simplified := denominators / numerators
	fmt.Println("problem 033:", simplified)
}

func problem_034() {
	var fact func(int) int
	fact = func(n int) int {
		if n < 2 {
			return 1
		}
		return n * fact(n-1)
	}

	facts := make(map[int]int)
	for i := range 10 {
		facts[i] = fact(i)
	}

	digit_sum := func(n int) int {
		sum := 0
		for n != 0 {
			sum += facts[n%10]
			n /= 10
		}
		return sum
	}

	total := 0
	for i := 3; i < 100000; i++ {
		if digit_sum(i) == i {
			total += i
		}
	}

	fmt.Println("problem 034:", total)
}

func problem_035() {
	prime_list := prime_sieve(1_000_000)
	prime_set := NewSet(prime_list...)

	count := 0
	checked := NewSet[int]()
	for _, p := range prime_list {
		if checked.Contains(p) {
			continue
		}
		digits := value_to_digits(p)
		circles := []int{}
		for _, v := range permute(digits) {
			circles = append(circles, digits_to_value(v))
		}
		slices.Sort(circles)
		circles = slices.Compact(circles)

		valid := true
		for _, c := range circles {
			if prime_set.Contains(c) {
				valid = false
				break
			}
			checked.Add(c)
		}
		if valid {
			fmt.Println(circles)
			count += len(circles)
		}
	}

	fmt.Println("problem 035:", count)
}

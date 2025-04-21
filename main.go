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

}

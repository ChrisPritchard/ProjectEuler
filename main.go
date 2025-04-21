package main

import (
	"fmt"
)

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()

	problem_031()
	problem_032()
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

	tester := func(set []int) int {
		sum := 0

		a := set[0]
		b := set[1]*1000 + set[2]*100 + set[3]*10 + set[4]
		c := set[5]*1000 + set[6]*100 + set[7]*10 + set[8]
		if _, exists := seen[c]; !exists && a*b == c {
			sum += c
			seen[c] = struct{}{}
		}

		a = set[0]*1000 + set[1]*100 + set[2]*10 + set[3]
		b = set[4]
		c = set[5]*1000 + set[6]*100 + set[7]*10 + set[8]
		if _, exists := seen[c]; !exists && a*b == c {
			sum += c
			seen[c] = struct{}{}
		}

		a = set[0]*10 + set[1]
		b = set[2]*100 + set[3]*10 + set[4]
		c = set[5]*1000 + set[6]*100 + set[7]*10 + set[8]
		if _, exists := seen[c]; !exists && a*b == c {
			sum += c
			seen[c] = struct{}{}
		}

		a = set[0]*100 + set[1]*10 + set[2]
		b = set[3]*10 + set[4]
		c = set[5]*1000 + set[6]*100 + set[7]*10 + set[8]
		if _, exists := seen[c]; !exists && a*b == c {
			sum += c
			seen[c] = struct{}{}
		}

		return sum
	}

	full_sum := 0
	for v := range all_permuts {
		full_sum += tester(all_permuts[v])
	}

	fmt.Println("problem 032:", full_sum)
}

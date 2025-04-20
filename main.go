package main

import "fmt"

func main() {
	Problems_001_010()
	Problems_011_020()
	Problems_021_030()

	problem_031()
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

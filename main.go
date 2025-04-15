package main

import "fmt"

func main() {
	Problems_001_010()
	Problems_011_020()

	problem_021()
}

func problem_021() {
	sum_divisors := func(n int) int {
		sum := 1
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		return sum
	}

	amicable_sum := 0
	added := make(map[int]struct{})
	for a := range 10_000 {
		if a < 2 {
			continue
		}
		b := sum_divisors(a)
		if a == b {
			continue
		}
		dB := sum_divisors(b)
		if dB == a {
			if _, exists := added[a]; !exists {
				added[a] = struct{}{}
				amicable_sum += a
			}
			if _, exists := added[b]; !exists {
				added[b] = struct{}{}
				amicable_sum += b
			}
		}
	}

	fmt.Println("problem 021:", amicable_sum)
}

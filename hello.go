package main

import (
	"fmt"
	"sort"
)

func main() {
	problem_001()
	problem_002()
	problem_003()
}

func problem_001() {
	sum := 0
	for i := 3; i < 1000; i += 3 {
		sum += i
	}
	for j := 5; j < 1000; j += 5 {
		if j%3 != 0 {
			sum += j
		}
	}
	fmt.Println("problem 001: ", sum)
}

func problem_002() {
	t1 := 1
	t2 := 2
	sum := 2
	for {
		n := t1 + t2
		if n > 4_000_000 {
			break
		}
		if n%2 == 0 {
			sum += n
		}
		t1 = t2
		t2 = n
	}
	fmt.Println("problem 002: ", sum)
}

func problem_003() {
	target := 600851475143
	primes := []int{}
	result := []int{}
	rem := target
	for n := 2; n < rem; n++ {
		is_valid := true
		for _, p := range primes {
			if n == p || n%p == 0 {
				is_valid = false
				break
			}
		}
		if !is_valid {
			continue
		}
		primes = append(primes, n)
		if rem%n == 0 {
			result = append(result, n)
			rem = target / n
		}
	}
	sort.Ints(result)
	fmt.Println("problem 003: ", result[len(result)-1])
}

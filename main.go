package main

import (
	"fmt"
	"sort"
)

func main() {
	problem_001()
	problem_002()
	problem_003()
	problem_004()
	problem_005()
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

	is_prime := func(n int) bool {
		for _, p := range primes {
			if n == p {
				return true
			}
			if n%p == 0 {
				return false
			}
		}
		primes = append(primes, n)
		return true
	}

	find_first_prime_divisor := func(v int) int {
		for n := 2; n < v; n++ {
			if v%n == 0 && is_prime(n) {
				return n
			}
		}
		return v
	}

	rem := target
	for {
		n := find_first_prime_divisor(rem)
		result = append(result, n)
		if n == rem {
			break
		}
		rem = rem / n
	}

	sort.Ints(result)
	fmt.Println("problem 003: ", result[len(result)-1])
}

func problem_004() {

	is_palindrome := func(n int) bool {
		if n == 0 {
			return false
		}
		v := n
		r := 0
		for v > 0 {
			r = r*10 + v%10
			v = v / 10
		}
		return n == r
	}

	var max *int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			v := i * j
			if max != nil && *max > v {
				continue
			}
			if is_palindrome(v) {
				max = &v
			}
		}
	}

	fmt.Println("problem 004: ", *max)
}

func problem_005() {
	greatest_common_divisor := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	least_common_multiple := func(a, b int) int {
		return (a / greatest_common_divisor(a, b)) * b
	}
	lcm := 1
	for i := 2; i <= 20; i++ {
		lcm = least_common_multiple(lcm, i)
	}
	fmt.Println("problem 005: ", lcm)
}

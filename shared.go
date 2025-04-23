package main

import "slices"

// calculate all primes in the range 1 to max using the prime sieve method
func prime_sieve(max int) []int {
	res := []int{}
	sieve := make([]bool, max)

	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}

		res = append(res, i)
		for j := i + i; j < len(sieve); j += i {
			sieve[j] = true
		}
	}

	return res
}

// simple integer power function
func pown(base int, power uint) int {
	if power == 0 {
		return 1
	}
	n := base
	for power > 1 {
		n *= base
		power--
	}
	return n
}

// reverses a slice and returns the result, without modifying the initial slice
func nrev[T comparable](slice []T) []T {
	n := make([]T, len(slice))
	for i := range slice {
		n[len(n)-1-i] = slice[i]
	}
	return n
}

// separates a number into its digits, e.g. 123 becomes 1,2,3
func value_to_digits(value int) []int {
	res := []int{}

	for value != 0 {
		res = append(res, value%10)
		value /= 10
	}

	return nrev(res)
}

// converts digits into a number, e.g. 1,2,3 becomes 123
func digits_to_value(digits []int) int {
	sum := 0
	for i, v := range nrev(digits) {
		sum += pown(10, uint(i)) * v
	}
	return sum
}

// returns all possible combinations of the input array
func permute(values []int) [][]int {

	var expander func([]int, []int) [][]int
	expander = func(current []int, rem []int) [][]int {
		results := [][]int{}
		for i, v := range rem {
			next := make([]int, len(current)+1)
			copy(next, current)
			next[len(next)-1] = v
			if len(rem) > 1 {
				new_rem := make([]int, len(rem))
				copy(new_rem, rem)
				new_rem = slices.Delete(new_rem, i, i+1)
				results = append(results, expander(next, new_rem)...)
			} else {
				results = append(results, next)
			}
		}
		return results
	}

	return expander([]int{}, values)
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	s.Add(items...)
	return s
}

func (s Set[T]) Add(items ...T) {
	for _, v := range items {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Remove(items ...T) {
	for _, v := range items {
		delete(s, v)
	}
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

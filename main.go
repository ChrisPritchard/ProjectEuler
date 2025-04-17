package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	Problems_001_010()
	Problems_011_020()

	problem_021()
	problem_022()
	problem_023()
	problem_024()
	problem_025()
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

func problem_022() {
	file, _ := os.Open("0022_names.txt")
	defer file.Close()

	words := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, p := range parts {
			words = append(words, strings.Trim(p, "\""))
		}
	}

	sort.Strings(words)

	sum := 0
	for i, v := range words {
		letter_sum := 0
		for _, c := range v {
			letter_sum += int(c) - int('A') + 1
		}
		sum += letter_sum * (i + 1)
	}

	fmt.Println("problem 022:", sum)
}

func problem_023() {

	memo := make(map[int]bool)

	abundant := func(n int) bool {
		if v, exists := memo[n]; exists {
			return v
		}
		if n < 12 {
			return false
		}
		sum := 1
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		res := sum > n
		memo[n] = res
		return res
	}

	not_a_sum := func(n int) bool {
		for i := 12; i <= n/2; i++ {
			if abundant(i) && abundant(n-i) {
				return false
			}
		}
		return true
	}

	sum := 0

	for n := range 28123 {
		if not_a_sum(n) {
			sum += n
		}
	}

	fmt.Println("problem 023:", sum)
}

func problem_024() {

	var fact func(n int) int
	fact = func(n int) int {
		if n < 2 {
			return 1
		}
		return n * fact(n-1)
	}

	available := []byte("0123456789")
	n := 1_000_000 - 1
	k := len(available)
	result := make([]byte, k)

	for i := range k {
		f := fact(k - 1 - i)
		index := n / f
		result[i] = available[index]
		available = slices.Delete(available, index, index+1)
		n -= index * f
	}

	fmt.Println("problem 024:", string(result))
}

func problem_025() {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	n := new(big.Int)
	i := 2

	for len(n.String()) < 1000 {
		n.Add(f1, f2)
		f1.Set(f2)
		f2.Set(n)
		i++
	}

	fmt.Println("problem 025:", i)
}

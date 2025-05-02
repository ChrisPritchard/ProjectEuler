package main

import (
	"fmt"
	"math/big"
	"os"
	"slices"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "all" {
		Problems_001_010()
		Problems_011_020()
		Problems_021_030()
		Problems_031_040()

		problem_041()
		problem_042()
		problem_043()
	}

	problem_044()
	problem_045()
	problem_046()
	problem_047()
	problem_048()
	problem_049()
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

func problem_044() {
	pent := func(i int) int {
		return (i * ((3 * i) - 1)) / 2
	}
	acc := NewSet[int]()
	i := 1
	for {
		n := pent(i)
		for j := 1; j < i; j++ {
			n2 := pent(j)
			n3 := n - n2
			if n3 == n2 || !acc.Contains(n3) {
				continue
			}
			diff := absi(n3 - n2)
			if acc.Contains(diff) {
				fmt.Println("problem 044:", diff)
				return
			}
		}
		acc.Add(n)
		i++
	}
}

func problem_045() {
	tri_number := func(n int) int {
		return (n * (n + 1)) / 2
	}
	pent_number := func(n int) int {
		return (n * (3*n - 1)) / 2
	}
	hent_number := func(n int) int {
		return n * (2*n - 1)
	}

	pents, hents := NewSet[int](), NewSet[int]()

	i := 286
	for {
		t := tri_number(i)
		p := pent_number(i)
		pents.Add(p)
		h := hent_number(i)
		hents.Add(h)
		i++

		if pents.Contains(t) && hents.Contains(t) {
			fmt.Println("problem 045:", t)
			return
		}
	}
}

func problem_046() {
	primes := prime_sieve(1_000_000)
	prime_set := NewSet(primes...)
	power_dubs := NewSet(2)
	for i := 2; i < 1_000; i++ {
		power_dubs.Add(i * i * 2)
	}

	for i := 9; i < 1_000_000; i += 2 {
		if prime_set.Contains(i) {
			continue
		}

		found := false
		for j := range primes {
			rem := i - primes[j]
			if power_dubs.Contains(rem) {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("problem 046:", i)
			return
		}
	}
}

func problem_047() {
	primes := prime_sieve(1_000_000)
	prime_set := NewSet(primes...)

	n := 4 // target consecutive number count
	m := 4 // target distinct prime factors

	count := 0
	i := 3
	for {
		if prime_set.Contains(i) {
			count = 0
			i++
			continue
		}

		rem := i
		factors := NewSet[int]()
		for rem > 1 {
			for j := range primes {
				p := primes[j]
				if rem%p == 0 {
					rem = rem / p
					factors.Add(p)
					break
				}
			}
		}
		if factors.Size() == m {
			count++
			if count == n {
				fmt.Println("problem 047:", i-(n-1))
				return
			}
		} else {
			count = 0
		}
		i++
	}
}

func problem_048() {
	n := big.NewInt(1)
	for i := 2; i <= 1000; i++ {
		o := big.NewInt(int64(i))
		p := big.NewInt(int64(i))
		for range i - 1 {
			o.Mul(o, p)
		}
		n.Add(n, o)
	}

	s := n.String()
	fmt.Println("problem 048:", s[len(s)-10:])
}

func problem_049() {
	in_range := func(i int) bool { return i >= 1000 && i <= 9999 }
	primes := NewSet(filter(prime_sieve(10000), in_range)...)

	for p := range primes {
		digits := value_to_digits(p)

		perms := transform(permute(digits), digits_to_value)
		perms = filter(filter(perms, in_range), primes.Contains)
		perms = NewSet(perms...).ToSlice()
		slices.Sort(perms)

		if len(perms) < 3 {
			continue
		}

		if perms[0] == 1487 {
			fmt.Println("debug")
		}

		fmt.Println(perms)

		diff := perms[1] - perms[0]
		valid := true
		for j := 1; j < len(perms)-1; j++ {
			if perms[j+1]-perms[j] != diff {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println(perms)
		}
	}
}

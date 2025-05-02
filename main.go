package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "all" {
		Problems_001_010()
		Problems_011_020()
		Problems_021_030()
		Problems_031_040()
		Problems_041_050()
	}

	problem_050()
}

func problem_050() {
	primes := prime_sieve(1_000_000)
	prime_set := NewSet(primes...)

	max := make(map[int]int)
	for i := range len(primes) - 1 {
		n := primes[i]
		for j := i + 1; j < len(primes); j++ {
			n += primes[j]
			if prime_set.Contains(n) {
				if m, exists := max[n]; !exists || m < j-i {
					max[n] = j - i
				}
			}
		}
	}

	mk := 0
	mv := 0
	for k, v := range max {
		if v > mv {
			mv = v
			mk = k
		}
	}

	fmt.Println("problem 050:", mk)
}

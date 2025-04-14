package main

import (
	"fmt"
	"sort"
)

func Problems_001_010() {
	problem_001()
	problem_002()
	problem_003()
	problem_004()
	problem_005()
	problem_006()
	problem_007()
	problem_008()
	problem_009()
	problem_010()
	problem_011()
	problem_012()
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
	fmt.Println("problem 001:", sum)
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
	fmt.Println("problem 002:", sum)
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
	fmt.Println("problem 003:", result[len(result)-1])
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

	fmt.Println("problem 004:", *max)
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
	fmt.Println("problem 005:", lcm)
}

func problem_006() {
	sum, sum_of_squares := 0, 0
	for i := 1; i <= 100; i++ {
		sum += i
		sum_of_squares += i * i
	}
	fmt.Println("problem 006:", sum*sum-sum_of_squares)
}

func problem_007() {
	primes := []int{2}

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

	for len(primes) <= 10000 {
		n := primes[len(primes)-1] + 1
		for {
			if is_prime(n) {
				break
			}
			n++
		}
	}

	fmt.Println("problem 007:", primes[len(primes)-1])
}

func problem_008() {
	n := `7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450`

	max := 0
	for i := range len(n) - 13 {
		chunk := n[i : i+13]
		product := 1
		for _, char := range chunk {
			value := int(char) - int(rune('0'))
			product *= value
		}
		if product > max {
			max = product
		}
	}

	fmt.Println("problem 008:", max)
}

func problem_009() {
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				if (a*a)+(b*b) == (c*c) && a+b+c == 1000 {
					fmt.Println("problem 009:", a*b*c)
					return
				}
			}
		}
	}
}

func problem_010() {
	primes := make([]bool, 2_000_000)

	sum := 0
	i := 2
	for {
		if !primes[i] {
			sum += i
			for j := i + i; j < len(primes); j += i {
				primes[j] = true
			}
		}
		i++
		if i >= len(primes) {
			break
		}
	}

	fmt.Println("problem 010:", sum)
}

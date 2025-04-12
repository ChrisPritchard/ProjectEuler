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

func problem_011() {
	a := [][]int{
		{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65},
		{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95},
		{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92},
		{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57},
		{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40},
		{04, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16},
		{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54},
		{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
	}

	max := 0

	for y := range len(a) {
		for x := range len(a[y]) {
			// across
			if x < len(a[y])-4 {
				s := a[y][x] * a[y][x+1] * a[y][x+2] * a[y][x+3]
				if s > max {
					max = s
				}
			}
			// down
			if y < len(a)-4 {
				s := a[y][x] * a[y+1][x] * a[y+2][x] * a[y+3][x]
				if s > max {
					max = s
				}
			}
			//tl-to-br
			if x < len(a[y])-4 && y < len(a)-4 {
				s := a[y][x] * a[y+1][x+1] * a[y+2][x+2] * a[y+3][x+3]
				if s > max {
					max = s
				}
			}

			//tr-to-bl
			if x >= 4 && y < len(a)-4 {
				s := a[y][x] * a[y+1][x-1] * a[y+2][x-2] * a[y+3][x-3]
				if s > max {
					max = s
				}
			}
		}
	}

	fmt.Println("problem 011:", max)
}

func problem_012() {

	triangle_number := func(n int) int {
		return (n * (n + 1)) / 2
	}

	divisor_count := func(n int) int {
		count := 1
		remaining := n

		if remaining%2 == 0 {
			exponent := 0
			for remaining%2 == 0 {
				exponent++
				remaining /= 2
			}
			count *= exponent + 1
		}

		divisor := 3
		for divisor*divisor <= remaining {
			if remaining%divisor == 0 {
				exponent := 0
				for remaining%divisor == 0 {
					exponent++
					remaining /= divisor
				}
				count *= exponent + 1
			}
			divisor += 2
		}

		if remaining > 1 {
			count *= 2
		}

		return count
	}

	i := 1
	for {
		if divisor_count(i) > 500 {
			fmt.Println("problem 012:", triangle_number(i))
			return
		}
		i++
	}

}

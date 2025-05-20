package main

import (
	"fmt"
	"math/big"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "all" {
		Problems_001_010()
		Problems_011_020()
		Problems_021_030()
		Problems_031_040()
		Problems_041_050()
	}

	problem_051()
	problem_052()
	problem_053()
	problem_054()
}

func problem_051() {
	// get a million primes
	// 8 primes could start at 0, 1 or 2 (3,4,5,6,7,8,9)
	// so go through primes, discarding any without the first three
	// when the first three are found, calculate positions and calculate candidates

	primes := prime_sieve(1_000_000)
	prime_set := NewSet(primes...)

	indexes := func(digits []int, looked_for int) []int {
		res := []int{}
		for i := range digits {
			if digits[i] == looked_for {
				res = append(res, i)
			}
		}
		return res
	}

	setter := func(digits []int, mask []int, replacement int) int {
		for i := range digits {
			if slices.Contains(mask, i) {
				digits[i] = replacement
			}
		}

		return digits_to_value(digits)
	}

	test_mask := func(digits []int, mask []int, start int) bool {
		count := 1
		for j := start; j <= 9; j++ {
			candidate := setter(digits, mask, j)
			if prime_set.Contains(candidate) {
				count++
			}
		}
		return count >= 8
	}

	for i := range primes {
		prime_digits := value_to_digits(primes[i])
		pos_0 := indexes(prime_digits, 0)
		if len(pos_0) > 0 {
			if test_mask(prime_digits, pos_0, 1) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
		pos_1 := indexes(prime_digits, 1)
		if len(pos_1) > 0 {
			if test_mask(prime_digits, pos_1, 2) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
		pos_2 := indexes(prime_digits, 2)
		if len(pos_2) > 0 {
			if test_mask(prime_digits, pos_2, 3) {
				fmt.Println("problem 051:", primes[i])
				return
			}
		}
	}
}

func problem_052() {

	ord := func(n int) []int {
		d := value_to_digits(n)
		slices.Sort(d)
		return d
	}

	n := 101

	for {
		d := ord(n)

		valid := slices.Compare(ord(2*n), d) == 0
		valid = valid && slices.Compare(ord(3*n), d) == 0
		valid = valid && slices.Compare(ord(4*n), d) == 0
		valid = valid && slices.Compare(ord(5*n), d) == 0
		valid = valid && slices.Compare(ord(6*n), d) == 0

		if !valid {
			n++
			continue
		}

		fmt.Println("problem 052:", n)
		return
	}
}

func problem_053() {
	memo := make(map[int]*big.Int)
	memo[0] = big.NewInt(0)
	var fact func(int) *big.Int
	fact = func(n int) *big.Int {
		if n <= 1 {
			return big.NewInt(1)
		}
		if v, exists := memo[n]; exists {
			return v
		}
		v := big.NewInt(int64(n))
		o := fact(n - 1)
		r := big.NewInt(0)
		r.Mul(v, o)
		memo[n] = r
		return r
	}

	nr := func(n, r int) *big.Int {
		num := fact(n)
		denom1 := fact(r)
		denom2 := fact(n - r)
		denom := big.NewInt(0)
		denom.Mul(denom1, denom2)
		res := big.NewInt(0)
		res.Div(num, denom)
		return res
	}

	count := 0
	target := big.NewInt(1_000_000)
	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			v := nr(i, j)
			if v.Cmp(target) > 0 {
				count++
			}
		}
	}

	fmt.Println("problem 053:", count)
}

func problem_054() {
	hands, _ := read_lines("./0054_poker.txt")

	cards := transform(strings.Split(hands[0], " "), func(token string) int {
		value := 0
		if unicode.IsDigit(rune(token[0])) {
			value = int(token[0]) - int('0')
		} else if token[0] == 'J' {
			value = 11
		} else if token[0] == 'Q' {
			value = 12
		} else if token[0] == 'K' {
			value = 13
		} else {
			value = 14
		}
		if token[1] == 'H' {
			value += 400
		} else if token[1] == 'D' {
			value += 300
		} else if token[1] == 'S' {
			value += 200
		} else {
			value += 100
		}
		return value
	})

	player_1 := cards[:5]

	is_royal_flush := func(cards []int) bool {
		same_suit := NewSet(transform(cards, func(n int) int { return n / 100 })...).Size() == 1
		if same_suit {
			slices.Sort(cards)
			return cards[0]%100 == 10
		} else {
			return false
		}
	}

	is_straight := func(cards []int) bool {
		same_suit := NewSet(transform(cards, func(n int) int { return n / 100 })...).Size() == 1
		if same_suit {
			slices.Sort(cards)
			return cards[4]-cards[0] == 4
		} else {
			return false
		}
	}

	// player_2 := cards[5:]

	// hand_type := func(cards []string) int {

	// }

	// player_1_win_count := 0
	// for i := range hands {

	// }
}

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
	problem_055()
	problem_056()
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
	type card struct {
		suit  byte
		value int
	}

	games, _ := read_lines("./0054_poker.txt")

	cards := func(hand []string) []card {
		return transform(hand, func(token string) card {
			value := 0

			if unicode.IsDigit(rune(token[0])) {
				value = int(token[0]) - int('0')
			} else if token[0] == 'T' {
				value = 10
			} else if token[0] == 'J' {
				value = 11
			} else if token[0] == 'Q' {
				value = 12
			} else if token[0] == 'K' {
				value = 13
			} else {
				value = 14
			}

			return card{value: value, suit: token[1]}
		})
	}

	hand_type := func(cards []card) int {
		cards = sort_by(cards, func(c card) int { return c.value })
		same_suit := NewSet(transform(cards, func(n card) byte { return n.suit })...).Size() == 1

		if same_suit && cards[0].value == 10 {
			return 10 // royal flush
		} else if same_suit && cards[4].value-cards[0].value == 4 {
			return 9 // straight flush
		}

		grouper := func(card card) int { return card.value }
		sortkey := func(pair KeyValue[int, []card]) int { return len(pair.Value)*100 + pair.Key }
		grouped := sort_by(to_slice(group_by(cards, grouper)), sortkey)

		if len(grouped) == 2 && len(grouped[1].Value) == 4 {
			return 8 // four of a kind
		}

		if len(grouped) == 2 && len(grouped[1].Value) == 3 {
			return 7 // full house
		}

		if same_suit {
			return 6 // flush
		}

		if len(grouped) == 5 && cards[4].value-cards[0].value == 4 {
			return 5 // straight
		}

		if len(grouped) == 3 && len(grouped[2].Value) == 3 {
			return 4 // three of a kind
		}

		if len(grouped) == 3 && len(grouped[2].Value) == 2 {
			return 3 // two pair
		}

		if len(grouped) == 4 {
			return 2 // one pair
		}

		return 1 // high card
	}

	first_higher_card := func(p1 []card, p2 []card) bool {
		grouper := func(card card) int { return card.value }
		sortkey := func(pair KeyValue[int, []card]) int { return len(pair.Value)*100 + pair.Key }

		grouped1 := sort_by(to_slice(group_by(p1, grouper)), sortkey)
		grouped2 := sort_by(to_slice(group_by(p2, grouper)), sortkey)

		for i := len(grouped1) - 1; i >= 0; i-- {
			if grouped1[i].Key > grouped2[i].Key {
				return true
			} else if grouped1[i].Key < grouped2[i].Key {
				return false
			}
		}

		panic("failed to calculate winner")
	}

	p1_wins := 0
	for i := range games {
		all_cards := strings.Split(games[i], " ")
		player_1 := cards(all_cards[:5])
		player_2 := cards(all_cards[5:])
		p1_hand := hand_type(player_1)
		p2_hand := hand_type(player_2)

		if p1_hand > p2_hand || (p1_hand == p2_hand && first_higher_card(player_1, player_2)) {
			p1_wins++
		}
	}

	fmt.Println("problem 054:", p1_wins)
}

func problem_055() {
	is_palindrome := func(num *big.Int) bool {
		digits := num.String()
		for i := range len(digits) / 2 {
			if digits[i] != digits[len(digits)-1-i] {
				return false
			}
		}
		return true
	}

	rev := func(i *big.Int) *big.Int {
		s := i.String()
		r := ""
		for _, v := range s {
			r = string(v) + r
		}
		res := big.NewInt(0)
		res.SetString(r, 10)
		return res
	}

	count := 0
	for i := range 10_000 {
		n := big.NewInt(int64(i))
		lychrel := true
		for range 50 {
			n.Add(n, rev(n))
			if is_palindrome(n) {
				lychrel = false
				break
			}
		}
		if lychrel {
			count++
		}
	}

	fmt.Println("problem 055:", count)
}

func problem_056() {
	calc_sum := func(a, b int) int {
		p := big.NewInt(int64(a))
		n := big.NewInt(100)
		for range b {
			n.Mul(n, p)
		}
		sum := 0
		for _, c := range n.String() {
			sum += int(c) - int('0')
		}
		return sum
	}
	max := 0
	for i := range 100 {
		for j := range 100 {
			v := calc_sum(i, j)
			if v > max {
				max = v
			}
		}
	}

	fmt.Println("problem 056:", max)
}

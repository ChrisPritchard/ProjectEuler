package main

func main() {
	problem_001()
	problem_002()
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
	println("problem 001: ", sum)
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
	println("problem 002: ", sum)
}

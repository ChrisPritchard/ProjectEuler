package main

func main() {
	problem_001()
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

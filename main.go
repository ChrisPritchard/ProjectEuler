package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Problems_001_010()
	Problems_011_020()

	problem_067()
}

func problem_067() {
	triangle := make([][]int, 0)

	file, _ := os.Open("0067_triangle.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, v := range strings.Fields(line) {
			n, _ := strconv.Atoi(v)
			row = append(row, n)
		}
		triangle = append(triangle, row)
	}

	max := triangle_counter_generic(triangle)
	fmt.Println("problem 067:", max)
}

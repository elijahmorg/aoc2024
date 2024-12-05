package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	program := readInput()
	// fmt.Println(program)

	// fmt.Printf("answer is: %d\n", part1(program))
	longTime := time.Now()
	for range 10000 {
		// start := time.Now()
		// fmt.Printf("answer is: %d\n", part2(program))
		part2(program)
		// fmt.Println("time: ", time.Since(start))
	}
	fmt.Println("total time:", time.Since(longTime))
}

func part1(data string) int {
	candidates := strings.Split(data, "mul(")

	var sum int
	for _, candidate := range candidates {
		sum += parseCandidate(candidate)
	}

	return sum
}

func part2(data string) int {
	doblocks := strings.Split(data, "do")

	var sum int
	for _, doblock := range doblocks {
		if doblock[:3] == "n't" {
			continue
		}
		candidates := strings.Split(doblock, "mul(")

		for _, candidate := range candidates {
			sum += parseCandidate(candidate)
		}
	}

	return sum
}

func parseCandidate(candidate string) int {
	// mul(1,1)
	// anything longer is bad
	// too short
	if len(candidate) < 4 {
		return 0
	}
	if len(candidate) > 8 {
		// trim mul( + any end
		candidate = candidate[:8]
	}
	i := strings.Index(candidate, ")")
	if i == -1 {
		return 0
	}
	maybe := candidate[:i]
	if len(maybe) == len(candidate) {
		// we should've removed data
		return 0
	}
	parts := strings.Split(maybe, ",")
	// should be an x and y
	if len(parts) != 2 {
		return 0
	}
	xStr := parts[0]
	yStr := parts[1]
	// too long
	if len(xStr) > 3 || len(yStr) > 3 {
		return 0
	}

	// make sure no funny characters
	if !isNum(xStr) || !isNum(yStr) {
		return 0
	}
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return 0
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return 0
	}

	return x * y
}

func isNum(numMaybe string) bool {
	for _, v := range numMaybe {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func isSameSign(a, b int) bool {
	if a > 0 && b > 0 {
		return true
	}
	if a < 0 && b < 0 {
		return true
	}
	return false
}

func readInput() string {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// lines := strings.Split(string(dat), "\n")
	//
	// for _, line := range lines {
	// 	var report []int
	// 	for _, v := range strings.Split(line, " ") {
	// 		num, _ := strconv.Atoi(v) // str number
	// 		report = append(report, num)
	// 	}
	// 	if len(report) == 1 {
	// 		continue
	// 	}
	// 	reports = append(reports, report)
	// }
	return string(dat)
}

func testCases() string {
	test := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	return test
}

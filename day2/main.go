package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reports := readInput()
	// reports := testCases()

	fmt.Printf("answer is: %d\n", part1(reports))
}

func part1(reports [][]int) int {
	sumSafe := 0
	fmt.Printf("len reports = %d\n", len(reports))
	for _, report := range reports {
		if isSafe(report) {
			sumSafe++
		}
	}
	return sumSafe
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	if !isGoodLoop(report) {
		// gave up and reverse this
		slices.Reverse(report)
		if !isGoodLoop(report) {
			return false
		}

	}

	return true // return early for false
}

func isGoodLoop(report []int) bool {
	badLevel := false
	neg := -1
	if (report[1] - report[0]) > 0 { // is inc or dec
		neg = 1
	}

	for i := 0; i < len(report)-1; i++ {
		if isBad(report[i], report[i+1], neg) {
			if !badLevel {
				badLevel = true
				// peek ahead
				if len(report)-2 > i && isBad(report[i], report[i+2], neg) {
					if len(report)-2 > i && isBad(report[i+1], report[i+2], neg) {
						return false
					} else {
						i++
					}
				} else {
					if len(report)-2 > i {
						report = append(report[:i], report[i+1:]...)
					}
					continue
				}
			}
			return false
		}
	}
	return true
}

func isBad(a, b, neg int) bool {
	diff := (b - a)
	if !isSameSign(diff, neg) {
		return true
	}
	absDiff := diff * neg
	if absDiff > 3 || absDiff < 1 {
		return true
	}
	return false
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

func readInput() [][]int {
	var reports [][]int

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		var report []int
		for _, v := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(v) // str number
			report = append(report, num)
		}
		if len(report) == 1 {
			continue
		}
		reports = append(reports, report)
	}

	return reports
}

func testCases() [][]int {
	// test := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{7, 6, 4, 2, 1},
	testStr := `7 6 4 2 1
7 8 11 14 16 18 19 25`

	lines := strings.Split(testStr, "\n")
	var reports [][]int

	for _, line := range lines {
		var report []int
		for _, v := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(v) // str number
			report = append(report, num)
		}
		if len(report) == 1 {
			continue
		}
		reports = append(reports, report)
	}

	return reports
}

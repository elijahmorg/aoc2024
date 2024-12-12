package main

import (
	"fmt"
	"strconv"
	"strings"
)

var visited [][]bool
var path [][]int
var obstacles = make(map[point]bool)

type point struct {
	x int
	y int
}

func main() {
	puzzleMap := readInput()
	// fmt.Println(puzzleMap)
	// fmt.Println(len(puzzleMap))
	// return
	// puzzleMap := testCases()
	// puzzle := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzleMap []string) int {
	for i := range 75 {
		puzzleMap = step(puzzleMap)
		fmt.Printf("step: %d len()=%d\n", i, len(puzzleMap))
	}
	return len(puzzleMap)
}

func applyRules(num string) []string {
	if num == "0" {
		return []string{"1"}
	}
	if len(num)%2 == 0 {
		right, _ := strconv.Atoi(num[:len(num)/2])
		left, _ := strconv.Atoi(num[len(num)/2:])
		return []string{strconv.Itoa(right), strconv.Itoa(left)}
	}

	numVal, _ := strconv.Atoi(num)

	return []string{strconv.Itoa(numVal * 2024)}
}

func step(puzzleMap []string) []string {
	for idx := 0; idx < len(puzzleMap); idx++ {
		val := puzzleMap[idx]
		newVals := applyRules(val)
		if len(newVals) == 2 {
			puzzleMap = append(puzzleMap[:idx+1], puzzleMap[idx:]...)
			puzzleMap[idx] = newVals[0]
			puzzleMap[idx+1] = newVals[1]
			idx++
			continue
		}
		puzzleMap[idx] = newVals[0]

	}

	// fmt.Println("print step")
	// fmt.Println(puzzleMap)

	return puzzleMap
}

func part2(puzzleMap []string) int {
	return 0
}

func readInput() []string {

	dat := "28591 78 0 3159881 4254 524155 598 1"
	parts := strings.Split(dat, " ")

	if len(parts[len(parts)-1]) == 0 {
		// get rid of last element
		parts = parts[:len(parts)-1]
	}

	return parts
}

func testCases() []string {
	test := `125 17`

	parts := strings.Split(string(test), " ")

	puzzleMap := parts

	return puzzleMap
}

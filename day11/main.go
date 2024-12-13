package main

import (
	"fmt"
	"math"
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
	// puzzleMap := testCases()
	// fmt.Println(puzzleMap)
	// fmt.Println(len(puzzleMap))
	// return
	// puzzle := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzleMap []string) int {
	return solveSlice(puzzleMap)
}

var (
	maxLevel = 75
)

func solveNum(num int, levels int, memo map[Key]int) int {
	if levels == maxLevel {
		return 1
	}
	key := Key{
		Num:   num,
		Level: levels,
	}

	if prev := memo[key]; prev != 0 {
		return prev
	}

	var total int
	if num == 0 {
		total = solveNum(1, levels+1, memo)

	} else if even, r, l := isEven(num); even {
		ra := solveNum(r, levels+1, memo)
		la := solveNum(l, levels+1, memo)
		total = ra + la
	} else {
		total = solveNum(num*2024, levels+1, memo)
	}

	memo[key] = total

	return total
}

func isEven(num int) (bool, int, int) {
	p := int(math.Log10(float64(num))) + 1
	if p%2 == 0 { // this seems off by one
		right := num % int(math.Pow10(p/2))
		left := num / int(math.Pow10(p/2))
		return true, right, left
	}

	return false, 0, 0
}

type Key struct {
	Level int
	Num   int
}

func solveSlice(puzzleMap []string) int {
	memo := make(map[Key]int)
	var count, levels int
	for idx := 0; idx < len(puzzleMap); idx++ {
		// fmt.Printf("idx=%d, count=%d", idx, count)
		val, _ := strconv.Atoi(puzzleMap[idx])
		count += solveNum(val, levels, memo)
	}

	// fmt.Println(memo)

	fmt.Println("answer: ", count)
	return count
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

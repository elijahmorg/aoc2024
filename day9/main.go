package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	puzzleMap := readInput()
	// puzzleMap := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzle []int) int {

	return solve(puzzle)
}

func part2(puzzle []int) int {
	return 0
}

func readInput() []int {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	parts := string(data)

	puzzleMap := make([]int, len(parts))
	for i := range puzzleMap {
		val, _ := strconv.Atoi(parts[i : i+1])
		puzzleMap[i] = val
	}

	return puzzleMap
}

func testCases() []int {
	test := `2333133121414131402`

	puzzleMap := make([]int, len(test))
	for i := range puzzleMap {
		val, _ := strconv.Atoi(test[i : i+1])
		puzzleMap[i] = val
	}

	return puzzleMap
}

func solve(puzzle []int) int {
	sum := calcSize(puzzle)
	mem := make([]int, sum)
	fillMemory(mem, puzzle)
	// fmt.Println(mem)
	sortMemory(mem)
	// fmt.Println(mem)
	return countMemory(mem)
}

func sortMemory(mem []int) {
	var last int

	last = len(mem) - 1
	for i, v := range mem {
		if v == -1 {
			for j := last; j > i; j-- {
				lastVal := mem[j]
				if lastVal != -1 {
					last = j                        // update
					mem[i], mem[j] = mem[j], mem[i] // swap
					break
				}
			}
		}
		if last == i {
			break
		}
	}
}

func countMemory(mem []int) int {
	var sum int
	for i, v := range mem {
		if v == -1 {
			break
		}
		sum += i * v
	}
	return sum
}

func fillMemory(mem []int, puzzle []int) {
	var memIndex int
	for i, val := range puzzle {
		for range val {
			// odds are free space
			if i%2 == 1 {
				mem[memIndex] = -1 // empty
			} else {
				// evens / 2 is file id
				mem[memIndex] = i / 2 // empty
			}
			memIndex = memIndex + 1
		}
	}

}

func calcSize(puzzle []int) int {
	var sum int
	for _, val := range puzzle {
		sum += val
	}

	return sum
}

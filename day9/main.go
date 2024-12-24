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
	return solve2(puzzle)
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

func solve2(puzzle []int) int {
	sum := calcSize(puzzle)
	mem := make([]int, sum)
	fillMemory(mem, puzzle)
	// fmt.Println(mem)
	sortMemory2(mem)
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

func sortMemory2(mem []int) {
	var last int
	last = len(mem) - 1
	// find file

	for last > 0 {
		var size int
		// pass in mem and last and then update last
		// last is the end of the file, so base is last - size
		last, size = findFile(mem, last)
		base := last - size + 1
		last = base - 1
		if last <= 0 || size == 0 {
			return
		}

		// fmt.Println(last, size)
		start := findSpace(mem, size) // look for space anywhere
		if start == -1 {
			// no space
			continue
		}
		if start > base {
			continue // can't swap to a later place
		}

		// swap
		for i := 0; i < size; i++ {
			mem[start+i], mem[base+i] = mem[base+i], mem[start+i]
		}
		// fmt.Println(mem)

		// we are done
	}
}

func findSpace(mem []int, l int) int {
	var start, cl int
	for i, v := range mem {
		if v == -1 {
			start = i
			cl++
			if cl == l {
				return start - cl + 1
			}
		} else {
			cl = 0
		}
	}
	return -1
}

// return start/index, size
func findFile(mem []int, index int) (int, int) {
	var start, size int

	for j := index; j > 0; j-- {
		lastVal := mem[j]
		// find file start
		if lastVal != -1 {
			start = j
			for ; lastVal == mem[j] && j > 0; j-- {
				size++
			}
			break
		}
	}

	return start, size
}

func countMemory(mem []int) int {
	var sum int
	for i, v := range mem {
		if v == -1 {
			continue
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

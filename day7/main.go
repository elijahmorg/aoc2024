package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// puzzleMap := readInput()
	puzzleMap := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzle []MathPuzzle) int {
	var ans int
	for _, p := range puzzle {
		if p.solve() {
			ans += p.Ans
		}
	}
	return ans
}

func part2(puzzle []MathPuzzle) int {
	var ans int
	for _, p := range puzzle {
		if p.solve2() {
			ans += p.Ans
		}
	}
	return ans
}

func readInput() []MathPuzzle {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	parts := strings.Split(string(dat), "\n")

	var puzzles []MathPuzzle
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		var p MathPuzzle
		var err error
		s1 := strings.Split(part, ":")
		p.Ans, err = strconv.Atoi(s1[0])
		if err != nil {
			panic("should be valid")
		}
		s2 := strings.Split(s1[1], " ")
		for _, val := range s2 {
			num, err := strconv.Atoi(val)
			if err != nil {
				// panic("should be valid")
				continue
			}
			p.vals = append(p.vals, num)
		}
		puzzles = append(puzzles, p)
	}

	return puzzles
}

func testCases() []MathPuzzle {
	test := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	parts := strings.Split(string(test), "\n")
	var puzzles []MathPuzzle
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		var p MathPuzzle
		var err error
		s1 := strings.Split(part, ":")
		p.Ans, err = strconv.Atoi(s1[0])
		if err != nil {
			panic("should be valid")
		}
		s2 := strings.Split(s1[1], " ")
		for _, val := range s2 {
			num, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			p.vals = append(p.vals, num)
		}
		puzzles = append(puzzles, p)
	}

	return puzzles
}

type MathPuzzle struct {
	Ans      int
	vals     []int
	Solvable bool
}

func (m *MathPuzzle) solve() bool {
	return solve(m.Ans, m.vals[0], m.vals[1:])
}

func (m *MathPuzzle) solve2() bool {
	return solve2(m.Ans, m.vals[0], m.vals[1:])
}

func solve(ans int, num int, nums []int) bool {
	if len(nums) == 0 {
		if ans == num {
			return true
		} else {
			return false
		}
	}
	if num > ans {
		// too big return early
		return false
	}
	c := nums[0]
	if solve(ans, num+c, nums[1:]) || solve(ans, num*c, nums[1:]) {
		return true
	}
	return false
}

func solve2(ans int, num int, nums []int) bool {
	if len(nums) == 0 {
		if ans == num {
			return true
		} else {
			return false
		}
	}
	if num > ans {
		// too big return early
		return false
	}
	c := nums[0]
	if solve2(ans, num+c, nums[1:]) || solve2(ans, num*c, nums[1:]) || solve2(ans, numCat(num, c), nums[1:]) {
		return true
	}
	return false
}

func numCat(num1, num2 int) int {
	a := strconv.Itoa(num1)
	b := strconv.Itoa(num2)

	c, _ := strconv.Atoi(a + b)
	return c
}

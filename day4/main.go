package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	puzzle := readInput()
	// puzzle := testCases()
	// fmt.Println(program)

	fmt.Printf("answer is: %d\n", part1(puzzle))
	fmt.Printf("answer is: %d\n", part2(puzzle))
}

func part1(puzzle string) int {
	plines := strings.Split(puzzle, "\n")
	if len(plines[len(plines)-1]) == 0 {
		plines = plines[:len(plines)-1]
	}
	var sum int
	for x := 0; x < len(plines); x++ {
		for y, v := range plines[x] {
			vChar := string(v)

			// start or end
			if vChar == "X" || vChar == "S" {
				if checkDiagonalRight(plines, x, y) {
					sum++
				}
				if checkDiagonalLeft(plines, x, y) {
					sum++
				}
				if checkHorizontal(plines, x, y) {
					sum++
				}
				if checkVertical(plines, x, y) {
					sum++
				}
			}
		}
	}

	return sum
}

func part2(puzzle string) int {
	plines := strings.Split(puzzle, "\n")
	if len(plines[len(plines)-1]) == 0 {
		plines = plines[:len(plines)-1]
	}
	var sum int
	for x := 0; x < len(plines); x++ {
		for y, v := range plines[x] {
			vChar := string(v)

			// start or end
			if vChar == "M" || vChar == "S" {
				if checkXMAS(plines, x, y) {
					sum++
				}
			}
		}
	}

	return sum
}

func checkXMAS(plines []string, x, y int) bool {
	// check enough left in horizontal space
	if len(plines[x]) <= y+2 {
		return false
	}
	// check if enough left in vertical space
	if len(plines) <= x+2 {
		return false
	}

	if checkXMASRight(plines, x, y) && checkXMASLeft(plines, x, y) {
		return true
	}

	return false
}

func checkXMASRight(plines []string, x, y int) bool {

	var diag []rune = make([]rune, 3)
	diag[0] = rune(plines[x][y])
	diag[1] = rune(plines[x+1][y+1])
	diag[2] = rune(plines[x+2][y+2])
	diagS := string(diag)
	if diagS == "MAS" || diagS == "SAM" {
		return true
	}

	return false
}

func checkXMASLeft(plines []string, x, y int) bool {
	y += 2

	var diag []rune = make([]rune, 3)
	diag[0] = rune(plines[x][y])
	diag[1] = rune(plines[x+1][y-1])
	diag[2] = rune(plines[x+2][y-2])
	diagS := string(diag)
	if diagS == "MAS" || diagS == "SAM" {
		return true
	}

	return false
}

func checkDiagonalRight(plines []string, x, y int) bool {
	// check enough left in horizontal space
	if len(plines[x]) <= y+3 {
		return false
	}
	// check if enough left in vertical space
	if len(plines) <= x+3 {
		return false
	}

	var diag []rune = make([]rune, 4)
	diag[0] = rune(plines[x][y])
	diag[1] = rune(plines[x+1][y+1])
	diag[2] = rune(plines[x+2][y+2])
	diag[3] = rune(plines[x+3][y+3])
	diagS := string(diag)
	if diagS == "XMAS" || diagS == "SAMX" {
		return true
	}

	return false
}

func checkDiagonalLeft(plines []string, x, y int) bool {
	// check enough left in horizontal space
	if y-3 < 0 {
		return false
	}
	// check if enough left in vertical space
	if len(plines) <= x+3 {
		return false
	}

	var diag []rune = make([]rune, 4)
	diag[0] = rune(plines[x][y])
	diag[1] = rune(plines[x+1][y-1])
	diag[2] = rune(plines[x+2][y-2])
	diag[3] = rune(plines[x+3][y-3])
	diagS := string(diag)
	if diagS == "XMAS" || diagS == "SAMX" {
		return true
	}

	return false
}

func checkHorizontal(plines []string, x, y int) bool {
	if len(plines[x]) <= y+3 {
		return false
	}
	sample := plines[x][y : y+4]
	if sample == "XMAS" || sample == "SAMX" {
		return true
	}

	return false
}

func checkVertical(plines []string, x, y int) bool {
	// check if enough left in vertical space
	if len(plines) <= x+3 {
		return false
	}

	var diag []rune = make([]rune, 4)
	diag[0] = rune(plines[x][y])
	diag[1] = rune(plines[x+1][y])
	diag[2] = rune(plines[x+2][y])
	diag[3] = rune(plines[x+3][y])
	diagS := string(diag)
	if diagS == "XMAS" || diagS == "SAMX" {
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
	test := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	return test
}

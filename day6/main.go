package main

import (
	"fmt"
	"os"
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
	// puzzle := testCases()

	visited = make([][]bool, len(puzzleMap))
	path = make([][]int, len(puzzleMap))
	for i := range visited {
		visited[i] = make([]bool, len(puzzleMap[0]))
		path[i] = make([]int, len(puzzleMap[0]))
	}

	fmt.Printf("answer is: %d\n", part1(puzzleMap))

	// printSolution(puzzleMap)
	visited = make([][]bool, len(puzzleMap))
	path = make([][]int, len(puzzleMap))
	for i := range visited {
		visited[i] = make([]bool, len(puzzleMap[0]))
		path[i] = make([]int, len(puzzleMap[0]))
	}
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
	// printSolution(puzzleMap)
}

func part1(puzzleMap []string) int {
	guard := findGuard(puzzleMap)
	// fmt.Println(guard)

	var sum int
	sum = 1
	for {
		var count, keepGoing bool
		guard, count, keepGoing = moveGuard(puzzleMap, guard)
		if !keepGoing {
			break
		}
		if count {
			// only count locations, not moves
			if val := visited[guard.y][guard.x]; !val {
				// sum++
				// mark visited spots
				visited[guard.y][guard.x] = true
			}
		}
	}
	for _, line := range visited {
		for _, count := range line {
			if count {
				sum++
			}
		}
	}
	return sum

}

const (
	North = iota
	East
	South
	West
)

type Guard struct {
	x, y      int
	direction int
}

// return updated guard or false if off screen
func moveGuard(puzzleMap []string, guard *Guard) (*Guard, bool, bool) {
	var count, keepGoing bool
	switch guard.direction {
	case North:
		guard, count, keepGoing = moveNorth(puzzleMap, guard, -1, -1)
	case East:
		guard, count, keepGoing = moveEast(puzzleMap, guard, -1, -1)
	case South:
		guard, count, keepGoing = moveSouth(puzzleMap, guard, -1, -1)
	case West:
		guard, count, keepGoing = moveWest(puzzleMap, guard, -1, -1)
	}

	return guard, count, keepGoing
}

// returns guard, moved, keepGoing
func moveNorth(puzzleMap []string, guard *Guard, x, y int) (*Guard, bool, bool) {
	// move off screen?
	if guard.y == 0 {
		return guard, true, false
	}
	// get above current location
	if x != -1 && y != -1 {
		if guard.x == x && guard.y-1 == y {
			// need to turn
			guard.direction = East
			return guard, false, true
		}
	}
	char := puzzleMap[guard.y-1][guard.x]
	if char == '#' {
		// need to turn
		guard.direction = East
		return guard, false, true
	}
	guard.y -= 1
	return guard, true, true
}

func moveEast(puzzleMap []string, guard *Guard, x, y int) (*Guard, bool, bool) {
	// move off screen?
	if guard.x == len(puzzleMap[0])-1 {
		return guard, true, false
	}
	// get above current location
	if x != -1 && y != -1 {
		if guard.x+1 == x && guard.y == y {
			// need to turn
			guard.direction = South
			return guard, false, true
		}
	}
	// get above current location
	char := puzzleMap[guard.y][guard.x+1]
	if char == '#' {
		// need to turn
		guard.direction = South
		return guard, false, true
	}
	guard.x += 1
	return guard, true, true
}

func moveSouth(puzzleMap []string, guard *Guard, x, y int) (*Guard, bool, bool) {
	// move off screen?
	if guard.y == len(puzzleMap)-1 {
		return guard, true, false
	}
	if x != -1 && y != -1 {
		if guard.x == x && guard.y+1 == y {
			// need to turn
			guard.direction = West
			return guard, false, true
		}
	}
	// get above current location
	char := puzzleMap[guard.y+1][guard.x]
	if char == '#' {
		// need to turn
		guard.direction = West
		return guard, false, true
	}
	guard.y += 1
	return guard, true, true
}

func moveWest(puzzleMap []string, guard *Guard, x, y int) (*Guard, bool, bool) {
	// move off screen?
	if guard.x == 0 {
		return guard, true, false
	}
	if x != -1 && y != -1 {
		if guard.x-1 == x && guard.y == y {
			// need to turn
			guard.direction = North
			return guard, false, true
		}
	}
	// get above current location
	char := puzzleMap[guard.y][guard.x-1]
	if char == '#' {
		// need to turn
		guard.direction = North
		return guard, false, true
	}
	guard.x -= 1
	return guard, true, true
}

var (
	guardStartX = -1
	guardStartY = -1
)

func findGuard(puzzleMap []string) *Guard {
	for y, line := range puzzleMap {
		x := strings.Index(line, "^")
		if x != -1 {
			guardStartX = x
			guardStartY = y
			return &Guard{x: x, y: y, direction: North}
		}
	}
	panic("no guard found")
}

func part2(puzzleMap []string) int {

	guard := findGuard(puzzleMap)

	var sum int
	for {
		var keepGoing bool
		guard, _, keepGoing = moveGuard(puzzleMap, guard)
		if !keepGoing {
			break
		}

		if ok, p := isSquare(puzzleMap, guard); ok {
			if _, ok := obstacles[*p]; !ok {
				obstacles[*p] = true
				sum++
			}
		}

		visited[guard.y][guard.x] = true
	}

	// printPath(puzzleMap, path)
	return sum
}

func printSolution(puzzleMap []string) {
	var lineLen int
	lineLen = len(puzzleMap[0])
	for y, line := range puzzleMap {
		if lineLen != len(line) {
			panic("all lines should be same length")
		}
		for x, char := range line {
			if visited[y][x] {
				if char == '#' {
					panic(fmt.Sprintf("can't visit # at %d,%d", x, y))
				}
				fmt.Printf("X")
				continue
			}
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
}
func printPath(puzzleMap []string, path [][]int) {

	var lineLen int
	lineLen = len(puzzleMap[0])
	for y, line := range puzzleMap {
		if lineLen != len(line) {
			panic("all lines should be same length")
		}
		for x, char := range line {
			if path[y][x] != 0 {
				if char == '#' {
					panic(fmt.Sprintf("can't visit # at %d,%d", x, y))
				}
				fmt.Printf("%d", path[y][x])
				continue
			}
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
}

func isSquare(puzzleMap []string, guard *Guard) (bool, *point) {
	var xStart, yStart, direction int
	// see if we come back
	xStart = guard.x
	yStart = guard.y
	direction = guard.direction
	currentPath := make(map[Guard]bool)

	defer func() {
		guard.y = yStart
		guard.x = xStart
		guard.direction = direction
	}()

	var obstacleX, obstacleY int
	if guard.direction == North {
		obstacleX = guard.x
		obstacleY = guard.y - 1
	}
	if guard.direction == East {
		obstacleX = guard.x + 1
		obstacleY = guard.y
	}
	if guard.direction == South {
		obstacleX = guard.x
		obstacleY = guard.y + 1
	}
	if guard.direction == West {
		obstacleX = guard.x - 1
		obstacleY = guard.y
	}
	if obstacleY == guardStartY && obstacleX == guardStartX {
		return false, nil
	}
	var keepGoing bool

	for {
		guard, _, keepGoing = moveGuardObstacle(puzzleMap, guard, obstacleX, obstacleY)
		p := Guard{x: guard.x, y: guard.y, direction: guard.direction}
		if !keepGoing {
			return false, nil
		}
		if _, ok := currentPath[p]; ok {
			path[obstacleY][obstacleX] = 9
			p := &point{obstacleX, obstacleY}
			return true, p
		}
		currentPath[p] = true
	}
}

// return updated guard or false if off screen
func moveGuardObstacle(puzzleMap []string, guard *Guard, x, y int) (*Guard, bool, bool) {
	var count, keepGoing bool
	switch guard.direction {
	case North:
		guard, count, keepGoing = moveNorth(puzzleMap, guard, x, y)
	case East:
		guard, count, keepGoing = moveEast(puzzleMap, guard, x, y)
	case South:
		guard, count, keepGoing = moveSouth(puzzleMap, guard, x, y)
	case West:
		guard, count, keepGoing = moveWest(puzzleMap, guard, x, y)
	}

	return guard, count, keepGoing
}

func readInput() []string {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	parts := strings.Split(string(dat), "\n")

	if len(parts[len(parts)-1]) == 0 {
		// get rid of last element
		parts = parts[:len(parts)-1]
	}

	return parts
}

func testCases() []string {
	// 	test := `....#.....
	// .........#
	// ..........
	// ..#.......
	// .......#..
	// ..........
	// .#..^.....
	// ........#.
	// #.........
	// ......#...`

	// 	test := `....##....
	// ........#.
	// ....^.....
	// .......#..
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........
	// ..........`

	test := `....##....
........#.
....^.#.#.
.......#..
..........
..........
..........
..........
..........
..........`

	parts := strings.Split(string(test), "\n")

	puzzleMap := parts

	return puzzleMap
}

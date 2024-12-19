package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := readInput()
	// puzzleInput := testCases()
	puzzleMap := makePuzzle(puzzleInput, 1024)

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %+v\n", part2(puzzleInput))
}

func makePuzzle(points []Point, tryLength int) [][]rune {
	var size int
	if len(points) == 25 {
		size = 6 + 1
	} else {
		size = 70 + 1
	}
	var puzzle [][]rune
	puzzle = make([][]rune, size)
	for i := range puzzle {
		puzzle[i] = make([]rune, size)
	}
	points = points[:tryLength]
	for _, point := range points {
		puzzle[point.y][point.x] = '#'
	}

	// printPuzzle(puzzle)
	return puzzle
}

func printPuzzle(puzzle [][]rune) {
	for y := range puzzle {
		for x := range puzzle {
			if puzzle[y][x] == '#' {
				fmt.Printf("#")

			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func part1(puzzle [][]rune) int {
	s := Point{x: 0, y: 0}
	e := Point{x: 70, y: 70}
	// if len(puzzle) == 70 {
	// 	e = Point{x: 70, y: 70}
	// }

	return solve(puzzle, s, e)
}

func part2(puzzleInput []Point) Point {
	// binary search
	var offset, half int
	half = (len(puzzleInput) - 1024) / 2

	offset = half

	fmt.Println("len()=", len(puzzleInput))

	for range len(puzzleInput) {
		puzzleMap := makePuzzle(puzzleInput, 1024+offset)
		sum := solve(puzzleMap, Point{0, 0, 0}, Point{70, 70, 0})
		fmt.Printf("ans: %d - %d\n", sum, 1024+offset)
		if sum == 0 {
			// 	return puzzleInput[1024+i]
			// }
			if half == 1 {
				for i := range 10 {
					puzzleMap := makePuzzle(puzzleInput, 1024+offset-i)
					sum = solve(puzzleMap, Point{0, 0, 0}, Point{70, 70, 0})
					if sum != 0 {
						fmt.Printf("ans: %d\n", sum)
						fmt.Println(1024 + offset - i + 1)
						return puzzleInput[1024+offset-i]
					}
				}
			}
			// 500 + 250
			// 750 - 250/2
			half = half / 2
			offset = offset - half
		} else {
			half = (len(puzzleInput) - 1024 - offset) / 2
			offset += half
		}
	}

	return Point{}
}

func readInput() []Point {

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

	var points []Point
	for _, part := range parts {
		var p Point
		coords := strings.Split(part, ",")
		p.x, _ = strconv.Atoi(coords[0])
		p.y, _ = strconv.Atoi(coords[1])
		points = append(points, p)

	}

	return points
}

func testCases() []Point {
	// 	test := `###############
	// #.......#....E#
	// #.#.###.#.###.#
	// #.....#.#...#.#
	// #.###.#####.#.#
	// #.#.#.......#.#
	// #.#.#####.###.#
	// #...........#.#
	// ###.#.#####.#.#
	// #...#.....#.#.#
	// #.#.#.###.#.#.#
	// #.....#...#.#.#
	// #.###.#.#.#.#.#
	// #S..#.....#...#
	// ###############`
	test := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	parts := strings.Split(string(test), "\n")
	var points []Point
	for _, part := range parts {
		var p Point
		coords := strings.Split(part, ",")
		var err error
		p.x, err = strconv.Atoi(coords[0])
		if err != nil {
			continue
		}
		p.y, _ = strconv.Atoi(coords[1])
		points = append(points, p)

	}

	return points
}

func solve(puzzle [][]rune, start, end Point) int {
	var c Cursor
	c.puzzle.p = puzzle
	c.puzzle.Size.y = len(puzzle)
	c.puzzle.Size.x = len(puzzle[0])
	c.p = start
	c.d = E
	// need a third dimension for visited, because direction is part of this
	var visited map[Point]bool = make(map[Point]bool)

	var dist map[Point]DistPoint = make(map[Point]DistPoint)

	// distance at start is 0
	dist[start] = DistPoint{start, 0, N}
	// visited[start] = true

	sm := (c.puzzle.Size.y * c.puzzle.Size.x)
	for range sm {
		u := getMinDistance(visited, dist)

		c.p = u
		visited[u] = true
		// if u.x == 6 && u.y == 1 {
		// 	fmt.Println(c.Edges())
		// }
		// if u.x == 6 && u.y == 2 {
		// 	fmt.Println(c.Edges())
		// }

		for _, edge := range c.Edges() {
			if ok := visited[edge.P]; !ok {
				cd := dist[u]
				cde, found := dist[edge.P]
				if !found {
					// no cost our cost is lower
					cd.Point = edge.P
					cd.Dist = cd.Dist + edge.Cost
					dist[edge.P] = cd
					continue
				} else if cd.Dist+edge.Cost < cde.Dist {
					cd.Point = edge.P
					cd.Dist = cd.Dist + edge.Cost
					dist[edge.P] = cd
				}
			}
		}
	}

	// PrintDistMap(dist, c.puzzle.Size, c.puzzle.p)
	// min := math.MaxInt
	// for k := range W {
	// 	p := end
	// 	p.d = k
	// 	if val, ok := dist[p]; ok && val.Dist < min {
	// 		min = val.Dist
	// 	}
	// }

	return dist[end].Dist
}

func PrintDistMap(dist map[Point]DistPoint, size Point, puzzle [][]rune) {
	for y := range size.y {
		for x := range size.x {
			p := Point{y: y, x: x}
			if dist[p].Dist == 0 {
				if puzzle[y][x] == '.' {
					fmt.Printf("000000")
				}
				fmt.Printf("077777 ")
			} else {
				fmt.Printf("%06d ", dist[p].Dist)
			}
		}
		fmt.Printf("\n")
	}
}

func getMinDistance(visited map[Point]bool, dist map[Point]DistPoint) Point {
	min := math.MaxInt
	p := Point{}
	for k, v := range dist {
		if visit, ok := visited[k]; ok && visit {
			continue
		}
		if v.Dist < min {
			min = v.Dist
			p = k
		}
	}

	return p
}

// cost and direction
type Edge struct {
	Cost int
	d    int
	P    Point
}

func (c *Cursor) Edges() []Edge {
	var edges []Edge
	b := makeBlock(c.puzzle.p, c.puzzle.Size, c.p)
	for k, v := range b.Touching {
		var edge Edge
		if v != nil {
			char := c.puzzle.p[v.y][v.x]
			if char == '#' {
				continue
			}
			edge.Cost = 1
			edge.d = k
			edge.P = *v
			edges = append(edges, edge)
		}
	}
	return edges
}

func minEdge(edges []Edge, visited map[Point]bool) Edge {
	min := edges[0].Cost
	ans := -1
	for i, edge := range edges {
		if edge.Cost < min && !visited[edge.P] {
			min = edge.Cost
			ans = i
		}
	}
	return edges[ans]
}

const (
	N = iota
	E
	S
	W
)

type Cursor struct {
	p      Point
	d      int // direction facing
	puzzle Puzzle
}

func (c *Cursor) Move(dir int) bool {
	if dir == N {
		if np := North(c.p); safeCheck(c.puzzle.Size, np) {
			c.p = np
			return true
		}
	}
	if dir == E {
		if ep := East(c.p); safeCheck(c.puzzle.Size, ep) {
			c.p = ep
			return true
		}
	}
	if dir == S {
		if sp := South(c.p); safeCheck(c.puzzle.Size, sp) {
			c.p = sp
			return true
		}
	}
	if dir == W {
		if wp := West(c.p); safeCheck(c.puzzle.Size, wp) {
			c.p = wp
			return true
		}
	}

	return false
}

func (c *Cursor) Turn(clockwise bool) {
	if true {
		c.d = c.d + 1%4
		return
	}
	c.d = c.d + 3%4 // minus one
}

type DistPoint struct {
	Point
	Dist int
	d    int
}

type Point struct {
	x int
	y int
	d int
}

type Puzzle struct {
	Size Point
	p    [][]rune
}

type Block struct {
	Point    Point
	Touching [4]*Point
	Rune     rune
}

func makeBlock(puzzle [][]rune, size, p Point) Block {
	var b Block
	b.Point = p
	b.Rune = puzzle[p.y][p.x]
	if np := North(p); safeCheck(size, np) {
		b.Touching[N] = &np
	}
	if ep := East(p); safeCheck(size, ep) {
		b.Touching[E] = &ep
	}
	if sp := South(p); safeCheck(size, sp) {
		b.Touching[S] = &sp
	}
	if wp := West(p); safeCheck(size, wp) {
		b.Touching[W] = &wp
	}

	return b
}

func North(p Point) Point {
	return Point{y: p.y - 1, x: p.x}
}

func East(p Point) Point {
	return Point{y: p.y, x: p.x + 1}
}

func South(p Point) Point {
	return Point{y: p.y + 1, x: p.x}
}

func West(p Point) Point {
	return Point{y: p.y, x: p.x - 1}
}

func safeCheck(size, p Point) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	if p.x >= size.x || p.y >= size.y {
		return false
	}

	return true
}

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	puzzleMap := readInput()
	// puzzleMap := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	// fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func findStartEnd(puzzleMap [][]rune) (Point, Point) {
	var start, end Point
	for y := range puzzleMap {
		for x := range puzzleMap[0] {
			if puzzleMap[y][x] == 'S' {
				start.x = x
				start.y = y
			}
			if puzzleMap[y][x] == 'E' {
				end.x = x
				end.y = y
			}
		}
	}
	return start, end
}

func part1(puzzle [][]rune) int {
	s, e := findStartEnd(puzzle)

	return solve(puzzle, s, e)
}

func part2(puzzle []string) int {
	return 0
}

func readInput() [][]rune {

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

	var puzzleMap [][]rune
	for _, part := range parts {
		puzzleMap = append(puzzleMap, []rune(part))
	}

	return puzzleMap
}

func testCases() [][]rune {
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
	test := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

	parts := strings.Split(string(test), "\n")
	var puzzleMap [][]rune
	for _, part := range parts {
		puzzleMap = append(puzzleMap, []rune(part))
	}

	return puzzleMap
}

type PuzzlePlus struct {
	Solution [][]Point
}

func printSolution() {
}

func solve(puzzle [][]rune, start, end Point) int {
	var c Cursor
	c.puzzle.p = puzzle
	c.puzzle.Size.y = len(puzzle)
	c.puzzle.Size.x = len(puzzle[0])
	c.p = start
	c.d = E
	// need a third dimension for visited, because direction is part of this
	var visited map[Point][4]bool = make(map[Point][4]bool)

	var dist map[Point]DistPoint = make(map[Point]DistPoint)

	// var parent map[Point]Point = make(map[Point]Point) // map[Point]parentPoint

	// distance at start is 0
	dist[start] = DistPoint{start, 0, E}
	// visited[start] = true

	sm := (c.puzzle.Size.y * c.puzzle.Size.x)
	for range sm {
		u := getMinDistance(visited, dist)

		c.p = u
		c.d = dist[u].d
		v := visited[u]
		v[c.d] = true
		visited[u] = v

		for _, edge := range c.Edges() {
			if b, ok := visited[edge.P]; !ok || !b[edge.d] {
				cd := dist[u]
				cde, found := dist[edge.P]
				if !found {
					// no cost our cost is lower
					cd.d = edge.d
					cd.x = edge.P.x
					cd.y = edge.P.y
					cd.Dist = cd.Dist + edge.Cost
					dist[edge.P] = cd
					continue
				} else if cd.Dist+edge.Cost < cde.Dist {
					cd.d = edge.d
					cd.x = edge.P.x
					cd.y = edge.P.y
					cd.Dist = cd.Dist + edge.Cost
					dist[edge.P] = cd
				}
			}
		}
	}

	return dist[end].Dist
}

func getMinDistance(visited map[Point][4]bool, dist map[Point]DistPoint) Point {
	min := math.MaxInt
	p := Point{}
	for k, v := range dist {
		if visit, ok := visited[k]; ok && visit[v.d] {
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
			// if valid move
			if char == '#' {
				continue
			}
			if k == c.d {
				edge.Cost = 1
				edge.d = k
				edge.P = *v
				edges = append(edges, edge)
			} else if k == c.d+2%4 {
				// behind us
				continue
			} else {
				edge.Cost = 1001
				edge.d = k
				edge.P = *v
				edges = append(edges, edge)
			}
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

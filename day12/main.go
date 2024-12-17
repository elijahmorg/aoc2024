package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	puzzleMap := readInput()
	// puzzleMap := testCases()

	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzle []string) int {
	groups := make(GroupMap)
	walkPuzzle(puzzle, groups)
	// size := Point{y: len(puzzle), x: len(puzzle[0])}
	groupsSet := getGroupsSet(groups)
	// printGroups(size, groupsSet)
	// fmt.Println(len(groupsSet))
	// fmt.Println(size)
	var cost int
	for group := range groupsSet {
		cost += group.Price(puzzle)

	}
	return cost
}

func part2(puzzle []string) int {
	return 0
}

func readInput() []string {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	parts := strings.Split(string(dat), "\n")
	// dat := "28591 78 0 3159881 4254 524155 598 1"
	// parts := strings.Split(dat, " ")

	if len(parts[len(parts)-1]) == 0 {
		// get rid of last element
		parts = parts[:len(parts)-1]
	}

	return parts
}

func testCases() []string {
	test := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	// 	test := `OOOOO
	// OXOXO
	// OOOOO
	// OXOXO
	// OOOOO`

	parts := strings.Split(string(test), "\n")

	puzzleMap := parts

	return puzzleMap
}

const (
	N = iota
	E
	S
	W
)

type Point struct {
	x int
	y int
}

type Block struct {
	Point    Point
	Touching [4]*Point // points of the same rune - index using NESW
	Rune     rune
}

type Group struct {
	Type   rune
	blocks map[Point]Block
}

func (g *Group) Price(puzzle []string) int {
	var perimeter int
	for _, b := range g.blocks {
		for _, t := range b.Touching {
			if t == nil {
				perimeter++
				continue
			}
			var nr rune
			nr = rune(puzzle[t.y][t.x])
			if nr != b.Rune {
				perimeter++
			}
		}
	}

	var cost int
	cost = len(g.blocks) * perimeter

	return cost
}

func (g *Group) String() string {
	return fmt.Sprintf("rune=%c/%d len()=%d", g.Type, g.Type, len(g.blocks))
}

type GroupMap map[Point]*Group

// algorithm - look up and left for same group.
// if left is same char as up but different group, merge groups
//
// walk matrix left to right top to bottom
func walkPuzzle(puzzle []string, groups GroupMap) {
	var size Point
	size.y = len(puzzle)
	size.x = len(puzzle[0])
	for y := range size.y {
		for x := range size.x {
			checkBlock(puzzle, size, Point{x: x, y: y}, groups)
		}
	}
}

func getGroupsSet(groups GroupMap) map[*Group]bool {
	setGroups := make(map[*Group]bool)

	for _, v := range groups {
		setGroups[v] = true
	}
	return setGroups
}

// pass in single block and add it to groups and merge groups
func checkBlock(puzzle []string, size, p Point, groups GroupMap) {
	b := makeBlock(puzzle, size, p)

	var leftMatch *Group
	var groupMatch bool
	if left := b.Touching[W]; left != nil {
		lRune := rune(puzzle[left.y][left.x])
		// check if same rune as left
		if lRune == b.Rune {
			groupMatch = true
			// add to left group
			lGroup := groups[*left]
			lGroup.blocks[p] = b
			groups[p] = lGroup
			leftMatch = lGroup
		}
	}
	var up *Point
	if up = b.Touching[N]; up != nil {
		upRune := rune(puzzle[up.y][up.x])
		// check if same rune as up
		if upRune == b.Rune {
			groupMatch = true
			// add to up group
			upGroup := groups[*up]
			if leftMatch != nil && leftMatch != upGroup {
				mergeGroups(leftMatch, upGroup, groups)
			} else {
				upGroup.blocks[p] = b
				groups[p] = upGroup
			}
		}
	}

	// didn't match left or up
	// make new group
	if !groupMatch {
		var newGroup Group
		newGroup.Type = b.Rune
		newGroup.blocks = make(map[Point]Block)
		newGroup.blocks[p] = b
		groups[p] = &newGroup
	}
}

// merge groupA into groupB
func mergeGroups(groupA, groupB *Group, groups GroupMap) {
	for p, v := range groupA.blocks {
		groups[p] = groupB
		groupB.blocks[p] = v
	}

}

func makeBlock(puzzle []string, size, p Point) Block {
	var b Block
	b.Point = p
	b.Rune = rune(puzzle[p.y][p.x])
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

func printSize(group *Group) {
	fmt.Printf("len(group)=%d", len(group.blocks))
}

func printGroups(size Point, groupSet map[*Group]bool) {
	for group, _ := range groupSet {
		printGroup(size, *group)
	}
}

// size is X,Y max point
func printGroup(size Point, group Group) {
	for y := range size.y {
		for x := range size.x {
			var p Point
			p.x = x
			p.y = y
			if val, ok := group.blocks[p]; ok {
				fmt.Printf("%c", val.Rune)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
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

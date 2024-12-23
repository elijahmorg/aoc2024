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
	// fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzle []Connection) int {

	return solve(puzzle)
}

func part2(puzzle []string) int {
	return 0
}

func readInput() []Connection {

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

	var puzzleMap []Connection
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		c := strings.Split(part, "-")
		if len(c) != 2 {
			panic("should be 2")
		}
		con := Connection{A: Node{N: c[0]}, B: Node{N: c[1]}}
		puzzleMap = append(puzzleMap, con)
	}

	return puzzleMap
}

func testCases() []Connection {
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
	test := `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

	parts := strings.Split(string(test), "\n")
	var puzzleMap []Connection
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		c := strings.Split(part, "-")
		if len(c) != 2 {
			panic("should be 2")
		}
		con := Connection{A: Node{N: c[0]}, B: Node{N: c[1]}}
		puzzleMap = append(puzzleMap, con)
	}

	return puzzleMap
}

type Connection struct {
	A Node
	B Node
}

type Node struct {
	N string
}

type NodeSet [3]Node

func (n *NodeSet) isT() bool {
	for _, v := range n {
		if v.N[0] == 't' {
			return true
		}
	}
	return false
}

/*

a-b
c-b
d-b
a-c
d-c
d-a

set[a] b, c , d
set[b] a, c, d
set[d] a, c, b

check(b, [c,d]) []node
for check range []
b conn to [c,d]
append []node
return ans // each node connected to a that is also connected to other nodes is a set of 3

so a-b-c, a-b-d are all sets
// mark A as done?
// we don't want to inspect b and recreate a-b-c we want unique


check(c, [d,b])


doneMap[a] = true
*/

type Puzzle struct {
	DoneMap map[Node]bool
	Sets    []NodeSet
	Conns   map[Node][]Node
}

func solve(puzzle []Connection) int {
	doneMap := make(map[Node]bool) // doneMap[ab] = true
	// count := 0
	var sets []NodeSet
	conns := getConnections(puzzle)
	p := Puzzle{DoneMap: doneMap, Conns: conns, Sets: sets}

	for a, an := range conns {
		if done := doneMap[a]; done {
			continue
		}
		for _, b := range an {
			sets = append(sets, checkConnections(a, b, an, p)...)
		}
		doneMap[a] = true
	}

	unique := make(map[NodeSet]bool)
	for _, v := range sets {
		if v.isT() {
			unique[v] = true
		}
	}

	// fmt.Println(conns, sets, count, doneMap, p)
	return len(unique)
}

// check if b is connected to other nodes also connected to a
func checkConnections(a, b Node, nodes []Node, p Puzzle) []NodeSet {
	// a is root node, b is check node
	var sets []NodeSet
	for _, n := range nodes {
		if n == b {
			// itself
			continue
		}
		bNodes := p.Conns[b]
		for _, k := range bNodes {
			if k == a || k == b {
				continue
			}
			if n == k {
				var set NodeSet
				set[0] = a
				set[1] = b
				set[2] = n
				set = sort(set)

				sets = append(sets, set)
			}
		}
	}

	return sets
}

func sort(set NodeSet) NodeSet {
	// if (el1 > el2) Swap(el1,el2)
	// if (el2 > el3) Swap(el2,el3)
	// if (el1 > el2) Swap(el1,el2)
	if string(set[0].N) > string(set[1].N) {
		set[1], set[0] = set[0], set[1]
	}
	if string(set[1].N) > string(set[2].N) {
		set[2], set[1] = set[1], set[2]
	}
	if string(set[0].N) > string(set[1].N) {
		set[1], set[0] = set[0], set[1]
	}
	return set
}

func getConnections(puzzle []Connection) map[Node][]Node {
	conns := make(map[Node][]Node)
	for _, v := range puzzle {
		addIfNot(v.A, v.B, conns) // add b to a
		addIfNot(v.B, v.A, conns) // add a to b
	}
	return conns
}

// add B to A
func addIfNot(A Node, B Node, conns map[Node][]Node) {
	if nodes, ok := conns[A]; ok {
		for _, v := range nodes {
			if v.N == B.N {
				return
			}
		}
		nodes = append(nodes, B)
		conns[A] = nodes // reassign just to be safe
	} else {
		conns[A] = []Node{B} // nothing there, just assign it
	}
}

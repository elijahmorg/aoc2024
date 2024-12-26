package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	puzzleMap := readInput()
	// puzzleMap := testCases()

	// fmt.Println(puzzleMap)
	fmt.Printf("answer is: %d\n", part1(puzzleMap))
	// fmt.Printf("answer is: %d\n", part2(puzzleMap))
}

func part1(puzzle Puzzle) int {
	return solve(puzzle)
}

func part2(puzzle []string) int {
	return 0
}

func readInput() Puzzle {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return Puzzle{}
	}

	p := puzzleParser(string(dat))

	return p
}

func testCases() Puzzle {
	test := `x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`

	p := puzzleParser(test)
	return p
}

func puzzleParser(input string) Puzzle {
	var p Puzzle
	p.Wires = make(map[string]*Wire)
	p.Gates = make(map[string]*Gate)

	parts := strings.Split(input, "\n\n")
	wires := strings.Split(parts[0], "\n")
	gates := strings.Split(parts[1], "\n")

	for _, wire := range wires {
		var w Wire
		_, err := fmt.Sscanf(wire, "%s %d", &w.Name, &w.Val)
		if err != nil {
			fmt.Println("about to panic:", wire)
			fmt.Println("error: ", err)
			panic("this shouldn't happen")
		}
		w.Name = w.Name[:3]
		w.Evaluated = true
		p.Wires[w.Name] = &w
	}

	if len(gates[len(gates)-1]) == 0 {
		gates = gates[:len(gates)-1]
	}

	for _, gate := range gates {
		var g Gate
		var a, b, c Wire
		var op string
		// tgd XOR rvg -> z12
		_, err := fmt.Sscanf(gate, "%s %s %s -> %s", &a.Name, &op, &b.Name, &c.Name)
		if err != nil {
			fmt.Println("about to panic:", gate)
			fmt.Println(err)
			panic("this shouldn't happen")
		}
		// if wire does not exist, create it.
		// null value is Evaluated = false and val = 0
		if _, ok := p.Wires[a.Name]; !ok {
			p.Wires[a.Name] = &a
		}
		if _, ok := p.Wires[b.Name]; !ok {
			p.Wires[b.Name] = &b
		}
		if _, ok := p.Wires[c.Name]; !ok {
			p.Wires[c.Name] = &c
		}
		g.A = p.Wires[a.Name]
		g.B = p.Wires[b.Name]
		g.Result = p.Wires[c.Name]

		if op == "OR" {
			g.Op = OR
		} else if op == "AND" {
			g.Op = AND
		} else if op == "XOR" {
			g.Op = XOR
		}

		p.Gates[a.Name+b.Name+c.Name] = &g
	}

	// find Z wires
	for n, v := range p.Wires {
		if n[0] == 'z' {
			p.Ans = append(p.Ans, v)
		}
	}
	sort.Slice(p.Ans, func(i, j int) bool {
		return p.Ans[i].Name < p.Ans[j].Name
	})
	for _, name := range p.Ans {
		fmt.Printf("%s ", name.Name)
	}
	fmt.Println()

	return p
}

type Wire struct {
	Name      string
	Val       int
	Evaluated bool
}

func (w *Wire) String() string {
	return fmt.Sprintf("%s, %t, %d", w.Name, w.Evaluated, w.Val)
}

const (
	AND = iota
	OR
	XOR
)

type Gate struct {
	A      *Wire
	B      *Wire
	Op     int
	Result *Wire
	Done   bool
}

func (g *Gate) Evaluate() bool {
	if g.Done {
		return true
	}
	if g.A == nil || g.B == nil {
		return false
	}
	if !g.A.Evaluated || !g.B.Evaluated {
		return false
	}
	g.Result.Val = Logic(g.Op, g.A, g.B)
	g.Result.Evaluated = true
	g.Done = true

	return false
}

func Logic(op int, A, B *Wire) int {
	if op == OR {
		return A.Val | B.Val
	}
	if op == AND {
		return A.Val & B.Val
	}
	if op == XOR {
		return A.Val ^ B.Val
	}

	fmt.Printf("should never get here")
	fmt.Printf("op: %d, A: %+v, B: %+v\n", op, A, B)
	panic("oops")
}

type Puzzle struct {
	Wires map[string]*Wire
	Gates map[string]*Gate // [A.Name+B.Name+Result.Name]*Gate
	Ans   []*Wire          // z wires
}

func ansDone(p Puzzle) bool {
	for _, ans := range p.Ans {
		if !ans.Evaluated {
			return false
		}
	}
	return true
}

func ansCalc(p Puzzle) int {
	var sum int
	for i, ans := range p.Ans {
		sum += ans.Val << i
	}
	return sum
}

func solve(p Puzzle) int {
	var numSolved int
	for _, w := range p.Wires {
		if w.Evaluated {
			numSolved++
		}
	}
	fmt.Println("numSolved: ", numSolved)
	for !ansDone(p) {
		fmt.Println("loop")
		for _, g := range p.Gates {
			g.Evaluate()
		}

		numSolved = 0
		for _, w := range p.Wires {
			if w.Evaluated {
				numSolved++
			}
		}
		fmt.Println("numSolved: ", numSolved)
	}
	return ansCalc(p)

}

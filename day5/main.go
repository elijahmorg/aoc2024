package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readInput()
	// rules, updates := testCases()
	// puzzle := testCases()
	// fmt.Println(program)

	fmt.Printf("answer is: %d\n", part1(rules, updates))
	fmt.Printf("answer is: %d\n", part2(rules, updates))
}

func part1(rules []*Rule, updates [][]int) int {
	ruleBook := make(map[int][]int)

	for _, rule := range rules {
		ruleBook[rule.a] = append(ruleBook[rule.a], rule.b)
	}

	var sum int
	for _, update := range updates {
		updateBook := make(map[int]int)
		for i, page := range update {
			updateBook[page] = i
		}

		if checkUpdate(update, ruleBook, updateBook) {
			if !isOdd(len(update)) {
				panic("should be odd")
			}
			val := update[len(update)/2] // for debugger
			sum += val
		}

	}

	return sum
}

func checkUpdate(update []int, ruleBook map[int][]int, updateBook map[int]int) bool {
	for i, page := range update {
		rules, ok := ruleBook[page]
		if !ok {
			// panic(fmt.Sprintf("bad rule? %d", page))
			continue
		}

		// check if rule page is before current pages location
		// it should be after
		for _, v := range rules {
			if found, ok := updateBook[v]; found < i && ok {
				return false
			}
		}

	}
	return true
}

// return indices so we can swap
func checkUpdateIndex(update []int, ruleBook map[int][]int, updateBook map[int]int) (bool, int, int) {
	for i, page := range update {
		rules, ok := ruleBook[page]
		if !ok {
			continue
		}

		// check if rule page is before current pages location
		// it should be after
		for _, v := range rules {
			if found, ok := updateBook[v]; found < i && ok {
				return false, i, found
			}
		}

	}
	return true, 0, 0
}

var recursionLevel int = 0

func fixUpdate(update []int, ruleBook map[int][]int, updateBook map[int]int) int {
	if recursionLevel > 200 {
		panic(fmt.Sprintf("too much cowbell: %+v", update))
	}
	ok, a, b := checkUpdateIndex(update, ruleBook, updateBook)
	if !ok {
		// swap place
		tmp := update[b]
		update[b] = update[a]
		update[a] = tmp
		recursionLevel += 1
		// overwrite previous
		updateBook = make(map[int]int)
		for i, page := range update {
			updateBook[page] = i
		}
		return fixUpdate(update, ruleBook, updateBook)
	}

	recursionLevel = 0
	return update[len(update)/2]
}

func isOdd(val int) bool {
	return val%2 == 1
}

func part2(rules []*Rule, updates [][]int) int {
	ruleBook := make(map[int][]int)

	for _, rule := range rules {
		ruleBook[rule.a] = append(ruleBook[rule.a], rule.b)
	}

	var sum int
	for _, update := range updates {
		updateBook := make(map[int]int)
		for i, page := range update {
			updateBook[page] = i
		}

		if !checkUpdate(update, ruleBook, updateBook) {
			val := fixUpdate(update, ruleBook, updateBook)
			sum += val

		}

	}

	return sum

}

type Rule struct {
	a int
	b int
}

func newRule(rule string) *Rule {
	var nr Rule
	parts := strings.Split(rule, "|")
	if len(parts) != 2 {
		panic("len should equal 2")
	}
	nr.a, _ = strconv.Atoi(parts[0])
	nr.b, _ = strconv.Atoi(parts[1])

	return &nr
}

func newUpdate(update string) []int {
	var updates []int
	for _, v := range strings.Split(update, ",") {
		n, _ := strconv.Atoi(v)
		updates = append(updates, n)
	}
	return updates
}

func readInput() ([]*Rule, [][]int) {

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	parts := strings.Split(string(dat), "\n\n")
	if len(parts) != 2 {
		panic("should be 2")
	}

	var rules []*Rule
	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			break
		}
		rules = append(rules, newRule(line))
	}

	var updates [][]int
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) == 0 {
			break
		}
		updates = append(updates, newUpdate(line))
	}

	return rules, updates
}

func testCases() ([]*Rule, [][]int) {
	test := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	parts := strings.Split(string(test), "\n\n")
	if len(parts) != 2 {
		panic("should be 2")
	}

	var rules []*Rule
	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) == 0 {
			break
		}
		rules = append(rules, newRule(line))
	}

	var updates [][]int
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) == 0 {
			break
		}
		updates = append(updates, newUpdate(line))
	}

	return rules, updates
}

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	right, left := readInput()

	slices.Sort(right)
	slices.Sort(left)

	sum := 0
	for i, k := range right {
		d := left[i] - k
		if d < 0 {
			d = d * -1
		}
		sum += d
	}
	fmt.Printf("answer: %d\n", sum)
	part2(right, left)
}

func part2(r, l []int) {
	count := make(map[int]int)

	var c int
	c = l[0]
	ccount := 0
	// list is already sorted
	for _, k := range r {
		if k == c {
			ccount++
		} else {
			count[k] = ccount
			ccount = 1
			c = k
		}
	}
	fmt.Println(count)

	var sim int
	for _, k := range l {
		val := count[k]
		sim += val * k
	}

	fmt.Printf("answer: %d\n", sim)
}

func readInput() ([]int, []int) {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	lines := strings.Split(string(dat), "\n")

	var r, l []int
	for _, line := range lines {
		numString := strings.Split(line, "   ")
		fmt.Println(numString)
		if len(numString) != 2 {
			break
		}
		rn, _ := strconv.Atoi(numString[0])
		ln, _ := strconv.Atoi(numString[1])

		r = append(r, rn)
		l = append(l, ln)
	}

	fmt.Println(r)
	fmt.Println(l)
	fmt.Println(len(l))

	return r, l
}

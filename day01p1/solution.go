package day01p1

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var first, second []int

	for _, ln := range lines {
		words := strings.Fields(ln)
		if len(words) != 2 {
			fmt.Printf("invalid input line: %q\n", ln)
			continue
		}

		f, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Printf("invalid number %q in line %q\n", words[0], ln)
			continue
		}

		s, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Printf("invalid number %q in line %q\n", words[1], ln)
			continue
		}

		first = append(first, f)
		second = append(second, s)
	}

	sort.Ints(first)
	sort.Ints(second)

	var diffSum int
	for i := range first {
		diffSum += abs(first[i] - second[i])
	}

	return diffSum
}

// math.Abs works on float64s
// Ryan Armstrong wrote up a really interesting post on how to implement an optimal solution:
// http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

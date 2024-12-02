package day01p2

import (
	"fmt"
	"io"
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

	freq := make(map[int]int)
	for _, n := range second {
		freq[n]++
	}

	var similarityScore int

	for _, n := range first {
		if count, exists := freq[n]; exists {
			similarityScore += n * count
		}
	}

	return similarityScore
}

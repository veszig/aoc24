package day03p1

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	mulRX := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var product int

	for _, l := range lines {
		for _, submatch := range mulRX.FindAllStringSubmatch(l, -1) {
			x, _ := strconv.Atoi(submatch[1])
			y, _ := strconv.Atoi(submatch[2])
			product += x * y
		}
	}

	return product
}

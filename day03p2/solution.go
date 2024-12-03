package day03p2

import (
	"io"
	"regexp"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	rx := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	mulRX := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var product int

	enabled := true
	for _, l := range lines {
		for _, match := range rx.FindAllString(l, -1) {
			switch match {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default: // mul
				if !enabled {
					continue
				}
				submatch := mulRX.FindStringSubmatch(match)
				x, _ := strconv.Atoi(submatch[1])
				y, _ := strconv.Atoi(submatch[2])
				product += x * y
			}
		}
	}

	return product
}

package day04p2

import (
	"io"

	"aoc/utils"
)

type coord struct {
	x int
	y int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var xmasCount int

	for i, line := range lines {
		for j, char := range line {
			if char == 'A' { // every time we find an A we'll check
				if isItDashXmas(&lines, i, j) {
					xmasCount += 1
				}
			}
		}
	}

	return xmasCount
}

func isItDashXmas(lines *[]string, i, j int) bool {
	topLeft := getRuneAt(lines, i-1, j-1)
	topRight := getRuneAt(lines, i-1, j+1)
	bottomLeft := getRuneAt(lines, i+1, j-1)
	bottomRight := getRuneAt(lines, i+1, j+1)

	return ((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) &&
		((topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M'))
}

func getRuneAt(lines *[]string, x, y int) rune {
	if x < 0 || x >= len(*lines) {
		return '.'
	}
	line := (*lines)[x]
	if y < 0 || y >= len(line) {
		return '.'
	}
	return rune(line[y])
}

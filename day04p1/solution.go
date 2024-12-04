package day04p1

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
			if char == 'X' { // every time we find an X we'll check in all directions
				xmasCount += countXmasInAllDirections(&lines, coord{x: i, y: j})
			}
		}
	}

	return xmasCount
}

func countXmasInAllDirections(lines *[]string, start coord) int {
	var count int

	// up
	if isItXmas(lines, start, coord{x: -1, y: -1}) {
		count += 1
	}
	if isItXmas(lines, start, coord{x: -1, y: 0}) {
		count += 1
	}
	if isItXmas(lines, start, coord{x: -1, y: 1}) {
		count += 1
	}
	// across
	if isItXmas(lines, start, coord{x: 0, y: -1}) {
		count += 1
	}
	if isItXmas(lines, start, coord{x: 0, y: 1}) {
		count += 1
	}
	// down
	if isItXmas(lines, start, coord{x: 1, y: -1}) {
		count += 1
	}
	if isItXmas(lines, start, coord{x: 1, y: 0}) {
		count += 1
	}
	if isItXmas(lines, start, coord{x: 1, y: 1}) {
		count += 1
	}

	return count
}

func isItXmas(lines *[]string, start coord, step coord) bool {
	return getRuneAt(lines, start.x, start.y) == 'X' &&
		getRuneAt(lines, start.x+step.x, start.y+step.y) == 'M' &&
		getRuneAt(lines, start.x+2*step.x, start.y+2*step.y) == 'A' &&
		getRuneAt(lines, start.x+3*step.x, start.y+3*step.y) == 'S'
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

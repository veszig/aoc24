package day10p2

import (
	"io"

	"aoc/utils"
)

type coord struct {
	row int
	col int
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var topoMap [][]int
	var trailheads []coord

	for row, line := range lines {
		var l []int
		for col, char := range line {
			i := int(char - '0')
			l = append(l, i)
			if i == 0 {
				trailheads = append(trailheads, coord{row: row, col: col})
			}
		}
		topoMap = append(topoMap, l)
	}

	sum := 0
	for _, th := range trailheads {
		sum += findTrails(&topoMap, th)
	}

	return sum
}

func findTrails(tMap *[][]int, pos coord) int {
	if (*tMap)[pos.row][pos.col] == 9 {
		return 1
	}

	sum := 0
	next := (*tMap)[pos.row][pos.col] + 1

	north := coord{row: pos.row - 1, col: pos.col}
	if north.row >= 0 && (*tMap)[north.row][north.col] == next {
		sum += findTrails(tMap, north)
	}

	east := coord{row: pos.row, col: pos.col + 1}
	if east.col < len((*tMap)[0]) && (*tMap)[east.row][east.col] == next {
		sum += findTrails(tMap, east)
	}

	south := coord{row: pos.row + 1, col: pos.col}
	if south.row < len(*tMap) && (*tMap)[south.row][south.col] == next {
		sum += findTrails(tMap, south)
	}

	west := coord{row: pos.row, col: pos.col - 1}
	if west.col >= 0 && (*tMap)[west.row][west.col] == next {
		sum += findTrails(tMap, west)
	}

	return sum
}

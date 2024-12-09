package day08p1

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

	var antennaMap [][]rune
	nodes := make(map[rune][]coord)
	for row, line := range lines {
		antennaMap = append(antennaMap, []rune(line))
		for col, char := range line {
			if char == '.' {
				continue
			}
			nodes[char] = append(nodes[char], coord{row: row, col: col})
		}
	}

	antinodeCandidates := make(map[coord]bool)
	for _, coords := range nodes {
		for n, antenna1 := range coords {
			for _, antenna2 := range coords[n+1:] {
				antinodeCandidates[coord{
					row: antenna1.row - (antenna2.row - antenna1.row),
					col: antenna1.col - (antenna2.col - antenna1.col),
				}] = true
				antinodeCandidates[coord{
					row: antenna2.row + (antenna2.row - antenna1.row),
					col: antenna2.col + (antenna2.col - antenna1.col),
				}] = true
			}
		}
	}

	antinodeCount := 0
	for coords := range antinodeCandidates {
		if coords.row < 0 || coords.col < 0 || coords.row >= len(antennaMap) ||
			coords.col >= len(antennaMap[0]) {
			continue
		}
		antinodeCount++
	}

	return antinodeCount
}

package day08p2

import (
	"io"

	"aoc/utils"
)

type coord struct {
	row int
	col int
}

func (c *coord) inBounds(maxRows, maxCols int) bool {
	if c.row < 0 || c.col < 0 || c.row >= maxRows || c.col >= maxCols {
		return false
	}
	return true
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

	maxRows := len(antennaMap)
	maxCols := len(antennaMap[0])

	antinodes := make(map[coord]bool)
	for _, coords := range nodes {
		for n, antenna1 := range coords {
			for _, antenna2 := range coords[n+1:] {
				antinodes[antenna1] = true
				antinodes[antenna2] = true
				i := 1
				for {
					candidate := coord{
						row: antenna1.row - i*(antenna2.row-antenna1.row),
						col: antenna1.col - i*(antenna2.col-antenna1.col),
					}
					if candidate.inBounds(maxRows, maxCols) {
						antinodes[candidate] = true
					} else {
						break
					}
					i++
				}
				j := 1
				for {
					candidate := coord{
						row: antenna2.row + j*(antenna2.row-antenna1.row),
						col: antenna2.col + j*(antenna2.col-antenna1.col),
					}
					if candidate.inBounds(maxRows, maxCols) {
						antinodes[candidate] = true
					} else {
						break
					}
					j++
				}
			}
		}
	}

	return len(antinodes)
}

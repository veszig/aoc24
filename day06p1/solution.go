package day06p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var guardMap [][]rune // strings are immutable, let's work with rune slices instead
	var guardX int        // guard position north -> south
	var guardY int        // guard position west -> east

	foundGuard := false
	for i, line := range lines {
		guardMap = append(guardMap, []rune(line))
		if !foundGuard {
			for j, char := range guardMap[i] {
				switch char {
				case 'v', '<', '^', '>':
					foundGuard = true
					guardX = i
					guardY = j
				default:
					continue
				}
			}
		}
	}

	var leftTheMap bool
	for {
		leftTheMap = moveGuard(&guardMap, &guardX, &guardY)
		if leftTheMap {
			break
		}
	}

	var xCount int
	for _, line := range guardMap {
		for _, char := range line {
			if char == 'X' {
				xCount += 1
			}
		}
	}

	return xCount
}

// returns true if the guard left the map
func moveGuard(guardMap *[][]rune, x, y *int) bool {
	switch (*guardMap)[*x][*y] {
	case '^':
		if *x == 0 { // guard is leaving the map
			(*guardMap)[*x][*y] = 'X'
			return true
		} else if (*guardMap)[(*x)-1][*y] == '#' { // guard rotates
			(*guardMap)[*x][*y] = '>'
			return false
		} else { // step north
			(*guardMap)[*x][*y] = 'X'
			(*guardMap)[(*x)-1][*y] = '^'
			*x -= 1
			return false
		}
	case '>':
		if *y == len((*guardMap)[*x])-1 { // guard is leaving the map
			(*guardMap)[*x][*y] = 'X'
			return true
		} else if (*guardMap)[*x][(*y)+1] == '#' { // guard rotates
			(*guardMap)[*x][*y] = 'v'
			return false
		} else { // step east
			(*guardMap)[*x][*y] = 'X'
			(*guardMap)[*x][(*y)+1] = '>'
			*y += 1
			return false
		}
	case 'v':
		if *x == len((*guardMap))-1 { // guard is leaving the map
			(*guardMap)[*x][*y] = 'X'
			return true
		} else if (*guardMap)[(*x)+1][*y] == '#' { // guard rotates
			(*guardMap)[*x][*y] = '<'
			return false
		} else { // step south
			(*guardMap)[*x][*y] = 'X'
			(*guardMap)[(*x)+1][*y] = 'v'
			*x += 1
			return false
		}
	case '<':
		if *y == 0 { // guard is leaving the map
			(*guardMap)[*x][*y] = 'X'
			return true
		} else if (*guardMap)[*x][(*y)-1] == '#' { // guard rotates
			(*guardMap)[*x][*y] = '^'
			return false
		} else { // step west
			(*guardMap)[*x][*y] = 'X'
			(*guardMap)[*x][(*y)-1] = '<'
			*y -= 1
			return false
		}
	default:
		panic("wtf")
	}
}

package day06p2

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

	// we'll change the position of the guard but later we'll also need the original position
	originalGuardX := guardX
	originalGuardY := guardY

	var leftTheMap bool
	for {
		leftTheMap = moveGuard(&guardMap, &guardX, &guardY)
		if leftTheMap {
			break
		}
	}

	// we'll keep track of the turns the guard does
	type turn struct {
		x              int
		y              int
		directionAfter rune
	}

	var oCount int
	for i, line := range guardMap {
	runeLoop:
		for j, char := range line {
			// an obstacle has to be somewhere where we would otherwise go
			if char == 'X' {
				// if the guard is standing here in the beginning, this isn't really a candidate
				if i == originalGuardX && j == originalGuardY {
					continue
				}

				// let's start with a fresh map
				var obstacleMap [][]rune
				for _, line := range lines {
					obstacleMap = append(obstacleMap, []rune(line))
				}

				// let's reset the guard position as well
				guardX = originalGuardX
				guardY = originalGuardY

				// put an obstacle in the place of the current X
				obstacleMap[i][j] = '#'

				// let's keep track of the turns encountered so far
				var turns []turn
				var x, y int // we'll keep track of the coordinates before the step here
				var leftTheMap bool
				for {
					x = guardX
					y = guardY
					leftTheMap = moveGuard(&obstacleMap, &guardX, &guardY)
					if leftTheMap {
						break // the obstacle didn't work
					}
					if x == guardX && y == guardY { // this was a turn
						currentTurn := turn{
							x:              x,
							y:              y,
							directionAfter: obstacleMap[x][y],
						}
						for _, t := range turns {
							if t.x == currentTurn.x && t.y == currentTurn.y && t.directionAfter == currentTurn.directionAfter {
								// we found a loop!
								oCount += 1
								continue runeLoop
							}
						}
						turns = append(turns, currentTurn)
					}
				}
			}
		}
	}

	return oCount
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

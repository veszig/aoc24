package day09p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	line := lines[0]

	var disk []int

	id := 0
	for i, char := range line {
		size := int(char - '0')
		blocks := make([]int, size)
		for j := range blocks {
			if i%2 == 0 { // the current character's position is even, so it is a file
				blocks[j] = id
			} else { // we are looking at a space
				blocks[j] = -1 // -1 means the block is empty
			}
		}
		disk = append(disk, blocks...)
		if i%2 == 0 {
			id++
		}
	}

	noFreeSpaceBefore := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 { // we don't move spaces
			continue
		}
		for disk[noFreeSpaceBefore] != -1 {
			noFreeSpaceBefore++
		}
		if noFreeSpaceBefore >= i {
			break
		}
		disk[noFreeSpaceBefore] = disk[i]
		disk[i] = -1
	}

	checksum := 0

	for i := range disk {
		if disk[i] == -1 {
			break
		}
		checksum += i * disk[i]
	}

	return checksum
}

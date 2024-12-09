package day09p2

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
			} else {
				blocks[j] = -1 // -1 means the block is empty
			}
		}
		disk = append(disk, blocks...)
		if i%2 == 0 {
			id++
		}
	}

	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}
		size := fileSizeBackwards(&disk, i)
		i -= (size - 1)
		pos := findFirstSpace(&disk, size)
		if pos == -1 || pos >= i {
			continue
		}
		for j := 0; j < size; j++ {
			disk[pos+j] = disk[i+j]
			disk[i+j] = -1
		}
	}

	checksum := 0

	for i := range disk {
		if disk[i] == -1 {
			continue
		}
		checksum += i * disk[i]
	}

	return checksum
}

func fileSizeBackwards(disk *[]int, position int) int {
	id := (*disk)[position]
	size := 0
	for position >= size && (*disk)[position-size] == id {
		size++
	}
	return size
}

func findFirstSpace(disk *[]int, size int) int {
	spaceSize := 0
	for i := range *disk {
		if (*disk)[i] == -1 {
			spaceSize++
		} else {
			spaceSize = 0
		}
		if spaceSize >= size {
			return i - spaceSize + 1
		}
	}
	return -1
}

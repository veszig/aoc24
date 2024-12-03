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

	doRX := regexp.MustCompile(`do(n't)?\(\)`)

	var product int

	enabled := true
	for _, l := range lines {
		doSubmatchIndexes := doRX.FindAllStringSubmatchIndex(l, -1)
		if doSubmatchIndexes == nil { // in this line there are no do() or don't() instructions
			if enabled {
				product += mulSum(l)
			}
		} else {
			if enabled { // we have to do the beginning of the line
				// TODO: this could be merged with the above empty line case
				product += mulSum(l[0:doSubmatchIndexes[0][0]])
			}
			for i, doSubmatchIdx := range doSubmatchIndexes {
				if doSubmatchIdx[2] == -1 { // we found a do()
					enabled = true
					// we have to process everything until the next instruction or the end of line
					if i+1 == len(doSubmatchIndexes) {
						// this is the last do() or don't() in the line, we have to process until the end of line
						product += mulSum(l[doSubmatchIdx[1]:])
					} else {
						// we have to process until the next do() or don't()
						product += mulSum(l[doSubmatchIdx[1]:doSubmatchIndexes[i+1][0]])
					}
				} else { // we found a don't()
					enabled = false
				}
			}
		}
	}

	return product
}

func mulSum(s string) int {
	mulRX := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var product int
	for _, submatch := range mulRX.FindAllStringSubmatch(s, -1) {
		x, _ := strconv.Atoi(submatch[1])
		y, _ := strconv.Atoi(submatch[2])
		product += x * y
	}
	return product
}

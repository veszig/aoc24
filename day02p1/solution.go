package day02p1

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines // each individual line is a report and each number is a level

	safeCount := 0
	for _, line := range lines {
		var report []int
		for _, l := range strings.Fields(line) {
			level, err := strconv.Atoi(l)
			if err != nil {
				fmt.Printf("invalid number %q in line %q", l, line)
				continue
			}
			report = append(report, level)
		}

		if isSafeReport(report) {
			safeCount++
		}
	}
	return safeCount
}

func isSafeReport(report []int) bool {
	increases := report[0] < report[1]

	var diff int

	for i := 0; i < len(report)-1; i++ {
		first := report[i]
		second := report[i+1]
		if first < second {
			if !increases {
				return false
			}
			diff = second - first
		} else {
			if increases {
				return false
			}
			diff = first - second
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

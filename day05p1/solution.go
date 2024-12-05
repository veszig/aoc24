package day05p1

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	orderingRX := regexp.MustCompile(`^(\d+)\|(\d+)$`)
	updatePagesRX := regexp.MustCompile(`^\d+(,\d+)*$`)

	// we'll leave the page numbers as strings
	pageMustBeBefore := make(map[string][]string)

	var sum int

lineLoop:
	for _, line := range lines {
		if submatch := orderingRX.FindStringSubmatch(line); len(submatch) > 0 {
			firstPage := submatch[1]
			secondPage := submatch[2]
			pageMustBeBefore[firstPage] = append(pageMustBeBefore[firstPage], secondPage)

		} else if match := updatePagesRX.MatchString(line); match {
			var seen []string
			pages := strings.Split(line, ",")
			for _, page := range pages {
				if currentPageMustBeBefore, exists := pageMustBeBefore[page]; exists {
					if hasIntersection(currentPageMustBeBefore, seen) {
						// this line is incorrect, we can ignore it
						continue lineLoop
					}
				}
				seen = append(seen, page)
			}

			// if we get here, we found a correct line
			middlePage := pages[len(pages)/2]
			middlePageNumber, _ := strconv.Atoi(middlePage)
			sum += middlePageNumber
		}
	}

	return sum
}

func hasIntersection(a, b []string) bool {
	m := make(map[string]struct{})
	for _, s := range a {
		m[s] = struct{}{}
	}

	for _, s := range b {
		if _, exists := m[s]; exists {
			return true
		}
	}

	return false
}

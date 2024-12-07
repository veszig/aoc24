package day07p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var sum int

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		target, _ := strconv.Atoi(parts[0])
		var arguments []int
		for _, num := range strings.Fields(strings.TrimSpace(parts[1])) {
			i, _ := strconv.Atoi(num)
			arguments = append(arguments, i)
		}
		if findMatchingOperators(target, arguments[0], arguments[1:]) {
			sum += target
		}
	}

	return sum
}

// returns true if we found a match and the argunets ran out
func findMatchingOperators(target int, acc int, remainingArgs []int) bool {
	if len(remainingArgs) > 0 {
		concatProduct, _ := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(remainingArgs[0]))
		return findMatchingOperators(target, acc+remainingArgs[0], remainingArgs[1:]) ||
			findMatchingOperators(target, acc*remainingArgs[0], remainingArgs[1:]) ||
			findMatchingOperators(target, concatProduct, remainingArgs[1:])
	} else if target == acc { // found one!
		return true
	} else {
		return false // :(
	}
}

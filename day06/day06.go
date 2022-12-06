package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 6

func compute(input string, length int) int {
	lastseen := make(map[rune]int)
	var i = 0
	res := 0
	input = strings.Trim(input, "\n")
	for {
		r := rune(input[i])
		if last, ok := lastseen[r]; ok && last >= res {
			res = last + 1
		}
		if len(input[res:i+1]) >= length {
			return res + length
		}
		lastseen[r] = i
		i++
	}
}

func part1(input string) interface{} {
	return compute(input, 4)
}

func part2(input string) interface{} {
	return compute(input, 14)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

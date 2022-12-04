package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 4

func contains(l1, h1, l2, h2 int) bool {
	return l1 >= l2 && h1 <= h2 || l2 >= l1 && h2 <= h1
}

func overlaps(l1, h1, l2, h2 int) bool {
	return l1 >= l2 && l1 <= h2 || h1 <= h2 && h1 >= l2 ||
		l2 >= l1 && l2 <= h1 || h2 <= h1 && h2 >= l1
}

func getRange(s string) (int, int) {
	bounds := strings.Split(s, "-")
	low, high := inputs.ParseDecInt(bounds[0]), inputs.ParseDecInt(bounds[1])
	return low, high
}

var score_p2 int

func part1(input string) interface{} {
	score := 0
	for _, line := range inputs.InputToStrList(input) {
		s := strings.Split(line, ",")
		l1, h1 := getRange(s[0])
		l2, h2 := getRange(s[1])
		if overlaps(l1, h1, l2, h2) {
			score_p2++
			if contains(l1, h1, l2, h2) {
				score++
			}
		}
	}
	return score
}

func part2(input string) interface{} {
	return score_p2
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

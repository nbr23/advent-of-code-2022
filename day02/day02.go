package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 2

func iWon(p1, p2 int) int {
	if p1 == p2 {
		return 3
	}
	// rock over cisors, cisors over paper, paper over rock
	if p1 == 0 && p2 == 2 || p1 == 2 && p2 == 1 || p1 == 1 && p2 == 0 {
		return 6
	}
	return 0
}

func part1(input string) interface{} {
	score := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}
		p1, p2 := int(line[0]-'A'), int(line[2]-'X')
		score += p2 + 1 + iWon(p2, p1)

	}
	return score
}

func part2(input string) interface{} {

	// rock needs paper, paper needs cisors, cisors need rock
	win := []int{1, 2, 0}
	// rock needs cisors, paper needsrock, cisor needs paper
	lose := []int{2, 0, 1}

	score := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}
		p1 := int(line[0] - 'A')
		switch line[2] {
		case 'X':
			score += 1 + lose[p1]
			break
		case 'Y':
			score += 1 + p1 + 3
			break
		case 'Z':
			score += 1 + win[p1] + 6
			break
		}

	}
	return score
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

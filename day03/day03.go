package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	"golang.org/x/exp/slices"
	//	"github.com/pkg/profile"
)

var day_num int = 3
var inputlist []string

func part1(input string) interface{} {
	inputlist = strings.Split(input, "\n")
	baditems := make([]int, 1000)
	baditems_count := 0
	score := 0
	for _, line := range inputlist {
		runes := []rune(line)
		size := len(runes) / 2
		for _, c := range runes[size : size*2] {
			if slices.Contains(runes[0:size], c) {
				if c >= 'A' && c <= 'Z' {
					baditems[baditems_count] = int(c) - int('A') + 27
				} else {
					baditems[baditems_count] = int(c) - int('a') + 1
				}
				score += baditems[baditems_count]
				baditems_count += 1
				break
			}
		}

	}
	return score
}

func part2(input string) interface{} {
	score := 0
	for i := 0; i < len(inputlist)-2; i += 3 {
		s1 := mapset.NewSet[rune]([]rune(inputlist[i])...)
		s2 := mapset.NewSet[rune]([]rune(inputlist[i+1])...)
		s3 := mapset.NewSet[rune]([]rune(inputlist[i+2])...)
		c, _ := s1.Intersect(s2).Intersect(s3).Pop()
		if c >= 'A' && c <= 'Z' {
			score += int(c) - int('A') + 27
		} else {
			score += int(c) - int('a') + 1
		}
	}
	return score
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

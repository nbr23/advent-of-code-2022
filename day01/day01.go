package main

import (
	"sort"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 1

var inputlist []string

func part1(input string) interface{} {
	inputlist = strings.Split(input, "\n")
	var current int = 0
	var max int = 0

	for _, s := range inputlist {
		if s == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += inputs.ParseDecInt(s)
		}
	}
	return max
}

func part2(input string) interface{} {
	var current int = 0
	var snacks []int = make([]int, 0)

	for _, s := range inputlist {
		if s == "" {
			i := sort.SearchInts(snacks, current)
			snacks = append(snacks, 0)
			copy(snacks[i+1:], snacks[i:])
			snacks[i] = current
			current = 0
		} else {
			current += inputs.ParseDecInt(s)
		}
	}
	l := len(snacks)
	return snacks[l-1] + snacks[l-2] + snacks[l-3]
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

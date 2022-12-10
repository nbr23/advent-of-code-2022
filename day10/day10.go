package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 10

func part1(input string) interface{} {
	X := 1
	res := 0
	waitcycle := false
	instructions := inputs.InputToStrList(input)
	cycle := 1
	for i := 0; i < len(instructions); {
		line := instructions[i]
		if line != "noop" {
			if waitcycle {
				X += inputs.ParseDecInt(strings.Split(line, " ")[1])
				i++
				waitcycle = false
			} else {
				waitcycle = true
			}
		} else {
			i++
		}
		cycle++
		if cycle == 20 || (cycle > 40 && cycle%40 == 20) {
			res += cycle * X
		}
	}
	return res
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 10

func draw(cycle, x int) rune {
	cycle = (cycle - 1) % 40
	if cycle >= x-1 && cycle <= x+1 {
		return '#'
	}
	return '.'
}

func part1(input string) interface{} {
	X := 1
	res := 0
	waitcycle := false
	instructions := inputs.InputToStrList(input)
	cycle := 1
	for i := 0; i < len(instructions); {
		fmt.Printf("%c", draw(cycle, X))
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
		if (cycle-1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	return res
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 25

var SNAFU = map[rune]int{
	'1': 1,
	'2': 2,
	'0': 0,
	'-': -1,
	'=': -2,
}

var UFANS = map[int]rune{
	1:  '1',
	2:  '2',
	0:  '0',
	-1: '-',
	-2: '=',
}

func power(a, p int) int {
	acc := 1
	for i := 0; i < p; i++ {
		acc *= a
	}
	return acc
}

func unsnafu(a int) string {
	s := ""
	for i := 0; ; i++ {
		fmt.Println(a)
		for k := -2; k <= 2; k++ {
			if (a-power(5, i)*k)%power(5, i+1) == 0 {
				s = fmt.Sprintf("%c%s", UFANS[k], s)
				a -= power(5, i) * k
				break
			}
		}
		if a == 0 {
			return s
		}
	}
}

func snafuParse(snafu string) int {
	res := 0
	pow := 0
	for i := len(snafu) - 1; i >= 0; i-- {
		res += power(5, pow) * SNAFU[rune(snafu[i])]
		pow++
	}
	return res
}

func part1(input string) interface{} {
	res := 0
	for _, l := range inputs.InputToStrList(input) {
		res += snafuParse(l)
	}
	return unsnafu(res)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

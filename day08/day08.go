package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 8

type Point struct {
	x, y int
}

func part1(input string) interface{} {
	trees := inputs.InputToIntMatrice(input)
	visible := 2*(len(trees)+len(trees[0])) - 4
	visimap := make(map[Point]bool)

	maxleft := make([]int, len(trees))
	maxright := make([]int, len(trees))
	maxtop := trees[0]
	maxbot := trees[len(trees)-1]
	for i := 0; i < len(trees); i++ {
		maxleft[i] = trees[i][0]
		maxright[i] = trees[i][len(trees[i])-1]
	}

	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[0])-1; x++ {
			curr := trees[y][x]
			if curr > maxtop[x] || curr > maxleft[y] {
				visimap[Point{x, y}] = true
			}
			if curr > maxtop[x] {
				maxtop[x] = curr
			}
			if curr > maxleft[y] {
				maxleft[y] = curr
			}
		}
	}

	for y := len(trees) - 2; y > 0; y-- {
		for x := len(trees[0]) - 2; x > 0; x-- {

			curr := trees[y][x]
			if curr > maxbot[x] || curr > maxright[y] {
				visimap[Point{x, y}] = true
			}
			if curr > maxbot[x] {
				maxbot[x] = curr
			}
			if curr > maxright[y] {
				maxright[y] = curr
			}
		}
	}
	return visible + len(visimap)
}

func scenicScore(trees [][]int, x, y int) int {
	self := trees[y][x]
	upscore, downscore, leftscore, rightscore := 0, 0, 0, 0

	// up
	for _y := y - 1; _y >= 0; _y-- {
		upscore++
		if trees[_y][x] >= self {
			break
		}
	}
	// down
	for _y := y + 1; _y < len(trees); _y++ {
		downscore++
		if trees[_y][x] >= self {
			break
		}
	}
	// left
	for _x := x + 1; _x < len(trees[0]); _x++ {
		leftscore++
		if trees[y][_x] >= self {
			break
		}
	}
	// right
	for _x := x - 1; _x >= 0; _x-- {
		rightscore++
		if trees[y][_x] >= self {
			break
		}
	}
	return upscore * downscore * leftscore * rightscore
}

func part2(input string) interface{} {
	trees := inputs.InputToIntMatrice(input)
	score := 0
	for y := len(trees) - 2; y > 0; y-- {
		for x := len(trees[0]) - 2; x > 0; x-- {
			s := scenicScore(trees, x, y)
			if s > score {
				score = s
			}
		}
	}
	return score
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

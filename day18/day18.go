package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 18

type Point struct {
	x, y, z int
}

var NEIGHBORS = []Point{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
	{-1, 0, 0},
	{0, -1, 0},
	{0, 0, -1},
}

func parseCubes(input string) int {
	res := 0
	cubes := make(map[Point]bool)
	for _, line := range inputs.InputToStrList(input) {
		coords := strings.Split(line, ",")
		cube := Point{
			inputs.ParseDecInt(coords[0]),
			inputs.ParseDecInt(coords[1]),
			inputs.ParseDecInt(coords[2]),
		}

		cubes[cube] = true
		res += 6
		for _, neighbor := range NEIGHBORS {
			if _, ok := cubes[Point{cube.x + neighbor.x, cube.y + neighbor.y, cube.z + neighbor.z}]; ok {
				res -= 2
			}
		}
	}
	return res
}

func part1(input string) interface{} {
	return parseCubes(input)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

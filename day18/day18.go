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

func countNeightbors(cubesMap map[Point]bool, c Point) int {
	count := 0
	for _, neighbor := range NEIGHBORS {
		if _, ok := cubesMap[Point{c.x + neighbor.x, c.y + neighbor.y, c.z + neighbor.z}]; ok {
			count++
		}
	}
	return count
}

type Bound struct {
	min, max int
}

func parseCubes(input string) (map[Point]bool, Bound, Bound, Bound) {
	boundX := Bound{100, 0} // min max
	boundY := Bound{100, 0}
	boundZ := Bound{100, 0}
	cubes := make(map[Point]bool)
	for _, line := range inputs.InputToStrList(input) {
		coords := strings.Split(line, ",")
		cube := Point{
			inputs.ParseDecInt(coords[0]),
			inputs.ParseDecInt(coords[1]),
			inputs.ParseDecInt(coords[2]),
		}
		boundX = Bound{utils.IntMin(boundX.min, cube.x), utils.IntMax(boundX.max, cube.x)}
		boundY = Bound{utils.IntMin(boundY.min, cube.y), utils.IntMax(boundY.max, cube.y)}
		boundZ = Bound{utils.IntMin(boundZ.min, cube.z), utils.IntMax(boundZ.max, cube.z)}

		cubes[cube] = true
	}
	return cubes, Bound{boundX.min - 1, boundX.max + 1}, Bound{boundY.min - 1, boundY.max + 1}, Bound{boundZ.min - 1, boundZ.max + 1}
}

func countExposed(cubes map[Point]bool) int {
	res := 0
	for cube := range cubes {
		res += 6
		for _, neighbor := range NEIGHBORS {
			if _, ok := cubes[Point{cube.x + neighbor.x, cube.y + neighbor.y, cube.z + neighbor.z}]; ok {
				res -= 1
			}
		}
	}
	return res
}

func part1(input string) interface{} {
	cubes, _, _, _ := parseCubes(input)
	return countExposed(cubes)
}

func addPoints(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y, p1.z + p2.z}
}

func steamSpread(cubes map[Point]bool, visited map[Point]bool, current Point, boundX, boundY, boundZ Bound, res *int) {
	for _, nb := range NEIGHBORS {
		neighbor := addPoints(current, nb)
		if neighbor.x > boundX.max || neighbor.x < boundX.min ||
			neighbor.y > boundY.max || neighbor.y < boundY.min ||
			neighbor.z > boundZ.max || neighbor.z < boundZ.min {
			continue
		}
		if _, ok := visited[neighbor]; ok {
			continue
		}
		if _, ok := cubes[neighbor]; ok {
			*res += 1
			continue
		}
		visited[neighbor] = true
		steamSpread(cubes, visited, neighbor, boundX, boundY, boundZ, res)
	}
}

func part2(input string) interface{} {
	res := 0
	cubes, boundX, boundY, boundZ := parseCubes(input)
	visited := make(map[Point]bool)
	steamSpread(cubes, visited, Point{boundX.min, boundY.min, boundZ.min}, boundX, boundY, boundZ, &res)
	return res
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

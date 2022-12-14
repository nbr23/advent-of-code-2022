package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 14

type Point struct {
	X int
	Y int
}

func parseRockPattern(input string, rockmap map[Point]bool) int {
	var prev *Point
	var maxy int
	for _, rock_s := range strings.Split(input, " -> ") {
		coords := strings.Split(rock_s, ",")
		rock := Point{inputs.ParseDecInt(coords[0]), inputs.ParseDecInt(coords[1])}

		if rock.Y > maxy {
			maxy = rock.Y
		}

		if prev == nil {
			prev = &rock
			continue
		}
		for x := prev.X; x <= rock.X; x++ {
			rockmap[Point{x, rock.Y}] = true
		}
		for x := rock.X; x <= prev.X; x++ {
			rockmap[Point{x, rock.Y}] = true
		}
		for y := prev.Y; y <= rock.Y; y++ {
			rockmap[Point{rock.X, y}] = true
		}
		for y := rock.Y; y <= prev.Y; y++ {
			rockmap[Point{rock.X, y}] = true
		}
		prev = &rock
	}
	return maxy
}

func addSand(rockmap map[Point]bool, floor int) bool {
	sand := Point{500, 0}
	for {
		if sand.Y >= floor {
			return false
		}
		if _, ok := rockmap[Point{sand.X, sand.Y + 1}]; !ok {
			sand.Y++
		} else if _, ok := rockmap[Point{sand.X - 1, sand.Y + 1}]; !ok {
			sand.Y++
			sand.X--
		} else if _, ok := rockmap[Point{sand.X + 1, sand.Y + 1}]; !ok {
			sand.Y++
			sand.X++
		} else {
			// The end
			rockmap[sand] = true
			return true
		}
	}
}

func part1(input string) interface{} {
	rockmap := make(map[Point]bool)
	var maxy int
	var res int
	for _, line := range inputs.InputToStrList(input) {
		maxy = parseRockPattern(line, rockmap)
	}
	for addSand(rockmap, maxy) {
		res++
	}
	return res
}

func addSand2(rockmap map[Point]bool, floor int) bool {
	sand := Point{500, 0}
	for {
		if sand.Y+1 == floor+2 {
			rockmap[sand] = true
			return true
		} else if _, ok := rockmap[Point{sand.X, sand.Y + 1}]; !ok {
			sand.Y++
		} else if _, ok := rockmap[Point{sand.X - 1, sand.Y + 1}]; !ok {
			sand.Y++
			sand.X--
		} else if _, ok := rockmap[Point{sand.X + 1, sand.Y + 1}]; !ok {
			sand.Y++
			sand.X++
		} else {
			if sand.Y == 0 {
				return false
			}
			rockmap[sand] = true
			return true
		}
	}
}

func part2(input string) interface{} {
	rockmap := make(map[Point]bool)
	var maxy int
	var res int
	for _, line := range inputs.InputToStrList(input) {
		maxy = parseRockPattern(line, rockmap)
	}
	for addSand2(rockmap, maxy) {
		res++
	}
	return res + 1
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

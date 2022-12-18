package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 17

type Point struct {
	X, Y int
}

type Rock struct {
	width   int
	height  int
	rockfun RockFunc
}

// -
func Rock_1(coord Point) []Point {
	return []Point{
		Point{coord.X, coord.Y},
		Point{coord.X + 1, coord.Y},
		Point{coord.X + 2, coord.Y},
		Point{coord.X + 3, coord.Y},
	}
}

var ROCK_1 = Rock{
	4,
	1,
	Rock_1,
}

// +
func Rock_2(coord Point) []Point {
	return []Point{
		Point{coord.X + 1, coord.Y},
		Point{coord.X, coord.Y + 1},
		Point{coord.X + 1, coord.Y + 1},
		Point{coord.X + 2, coord.Y + 1},
		Point{coord.X + 1, coord.Y + 2},
	}
}

var ROCK_2 = Rock{
	3,
	1,
	Rock_2,
}

// L backwards
func Rock_3(coord Point) []Point {
	return []Point{
		Point{coord.X + 2, coord.Y + 2},
		Point{coord.X + 2, coord.Y + 1},
		Point{coord.X, coord.Y},
		Point{coord.X + 1, coord.Y},
		Point{coord.X + 2, coord.Y},
	}
}

var ROCK_3 = Rock{
	3,
	1,
	Rock_3,
}

// l
func Rock_4(coord Point) []Point {
	return []Point{
		Point{coord.X, coord.Y},
		Point{coord.X, coord.Y + 1},
		Point{coord.X, coord.Y + 2},
		Point{coord.X, coord.Y + 3},
	}
}

var ROCK_4 = Rock{
	1,
	1,
	Rock_4,
}

// square
func Rock_5(coord Point) []Point {
	return []Point{
		Point{coord.X, coord.Y},
		Point{coord.X + 1, coord.Y},
		Point{coord.X, coord.Y + 1},
		Point{coord.X + 1, coord.Y + 1},
	}
}

var ROCK_5 = Rock{
	2,
	1,
	Rock_5,
}

type RockFunc func(coord Point) []Point

var ROCKS = []Rock{
	ROCK_1,
	ROCK_2,
	ROCK_3,
	ROCK_4,
	ROCK_5,
}

func canGoTo(playground map[Point]bool, current_rock *Rock, current_rock_coords Point) bool {
	for _, coords := range current_rock.rockfun(current_rock_coords) {
		if coords.X < 0 || coords.X >= 7 {
			return false
		}
		if _, ok := playground[coords]; ok {
			return false
		}
	}
	return true
}

func fillFloor(playground map[Point]bool) {
	for i := 0; i <= 7; i++ {
		playground[Point{i, 0}] = true
	}
}

func printMap(playground map[Point]bool, highest int, current Rock, current_pos Point) {
	for y := highest; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if _, ok := playground[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				done := false
				for _, pos := range current.rockfun(current_pos) {
					if pos.X == x && pos.Y == y {
						fmt.Print("@")
						done = true
						break
					}
				}
				if !done {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func dropRocks(input string, MAX int) int {
	playground := make(map[Point]bool)
	fillFloor(playground)
	rockindex := 0
	highest := 0
	var current_rock *Rock
	current_rock = nil
	var current_rock_coord Point // top left coordinates

	rock_count := MAX + 1

	for i := 0; ; i++ {
		if current_rock == nil {
			current_rock = &ROCKS[rockindex]
			current_rock_coord = Point{2, highest + current_rock.height + 3}
			rockindex = (rockindex + 1) % len(ROCKS)
			rock_count--
			if rock_count == 0 {
				fmt.Println("Done?")
				return highest
			}
		}

		c := input[i%len(input)]

		// Lateral move
		if c == '>' {
			if canGoTo(playground, current_rock, Point{current_rock_coord.X + 1, current_rock_coord.Y}) {
				current_rock_coord.X++
			} else {
			}
		} else {
			if canGoTo(playground, current_rock, Point{current_rock_coord.X - 1, current_rock_coord.Y}) {
				current_rock_coord.X--
			} else {
			}
		}

		// down
		if canGoTo(playground, current_rock, Point{current_rock_coord.X, current_rock_coord.Y - 1}) {
			current_rock_coord.Y--
		} else {
			for _, rock := range current_rock.rockfun(current_rock_coord) {
				highest = utils.IntMax(highest, rock.Y)
				playground[rock] = true
			}
			current_rock = nil
		}
	}
	return highest
}

func part1(input string) interface{} {

	return dropRocks(strings.Trim(input, "\n"), 2022)
}

func part2(input string) interface{} {
	return dropRocks(strings.Trim(input, "\n"), 1000000000000)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

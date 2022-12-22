package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 22

type point struct {
	x int
	y int
}

func loadMap(input []string) map[point]bool {
	monkeymap := make(map[point]bool)
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				monkeymap[point{x, y}] = true
			} else if input[y][x] == '.' {
				monkeymap[point{x, y}] = false
			}
		}
	}
	return monkeymap
}

type turtle struct {
	pos    point
	facing point
}

var FACING_NUM = map[point]int{
	point{1, 0}:  0,
	point{0, 1}:  1,
	point{-1, 0}: 2,
	point{0, -1}: 3,
}

func rotate(facing *point, angle int) {
	if facing.x != 0 {
		facing.y = facing.x * angle / -90
		facing.x = 0
	} else {
		facing.x = facing.y * angle / 90
		facing.y = 0
	}
	// 0,1 + 90 ->
}

func moveTurtle(monkeyMap map[point]bool, t *turtle, steps, maxx, maxy int) {
	for i := 0; i < steps; i++ {
		v, ok := monkeyMap[point{t.pos.x + t.facing.x, t.pos.y + t.facing.y}]
		if !ok { // out of map

			// loop horizontal positive
			if t.facing.x == 1 {
				newx := 0
				for {
					v, ok = monkeyMap[point{newx, t.pos.y}]
					if ok {
						if v {
							// we loop around but end up on a rock, so no move.
							return
						} else {
							t.pos.x = newx
							break
						}
					}
					newx++
				}
			}

			// loop horizontal negative
			if t.facing.x == -1 {
				newx := maxx
				for {
					v, ok = monkeyMap[point{newx, t.pos.y}]
					if ok {
						if v {
							// we loop around but end up on a rock, so no move.
							return
						} else {
							t.pos.x = newx
							break
						}
					}
					newx--
				}
			}

			// loop vertical positive
			if t.facing.y == 1 {
				newy := 0
				for {
					v, ok = monkeyMap[point{t.pos.x, newy}]
					if ok {
						if v {
							// we loop around but end up on a rock, so no move.
							return
						} else {
							t.pos.y = newy
							break
						}
					}
					newy++
				}
			}

			// loop vertical negative
			if t.facing.y == -1 {
				newy := maxy
				for {
					v, ok = monkeyMap[point{t.pos.x, newy}]
					if ok {
						if v {
							// we loop around but end up on a rock, so no move.
							return
						} else {
							t.pos.y = newy
							break
						}
					}
					newy--
				}
			}

		} else {
			if !v { // we can go
				t.pos.x += t.facing.x
				t.pos.y += t.facing.y
			} else {
				// we can just exit, we won't move any further
				return
			}
		}
	}
}

func getMaxX(input []string) int {
	res := 0
	for _, l := range input {
		res = utils.IntMax(res, len(l))
	}
	return res
}

// angles, steps
func splitDirectionsChunks(input string) ([]int, []int) {
	steps := make([]int, 0)
	angles := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == 'L' {
			angles = append(angles, 90)
		} else if input[i] == 'R' {
			angles = append(angles, -90)
		} else {
			k := i
			for ; k < len(input) && input[k] != 'L' && input[k] != 'R'; k++ {
			}
			steps = append(steps, inputs.ParseDecInt(input[i:k]))
			i = k - 1
		}
	}
	return angles, steps
}

func getFirstTileX(monkeyMap map[point]bool) int {
	x := 0
	for ; x < len(monkeyMap); x++ {
		v, ok := monkeyMap[point{x, 0}]
		if ok && !v {
			break
		}
	}
	return x
}

func printMap(monkeyMap map[point]bool, maxx, maxy int, t turtle) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			v, ok := monkeyMap[point{x, y}]
			if t.pos.x == x && t.pos.y == y {
				fmt.Print("@")
			} else if !ok {
				fmt.Print(" ")
			} else if !v {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func part1(input string) interface{} {
	parts := strings.Split(strings.Trim(input, "\n"), "\n\n")

	angles, steps := splitDirectionsChunks(parts[1])

	monkeyMap_s := strings.Split(parts[0], "\n")
	monkeyMap := loadMap(monkeyMap_s)

	t := turtle{
		point{getFirstTileX(monkeyMap), 0},
		point{1, 0},
	}

	maxx := getMaxX(monkeyMap_s)
	maxy := len(monkeyMap_s)

	for i, step := range steps {
		// printMap(monkeyMap, maxx, maxy, t)
		moveTurtle(monkeyMap, &t, step, maxx, maxy)
		if i < len(angles) {
			rotate(&t.facing, angles[i])
		}
	}
	t.pos.y++
	t.pos.x++
	return (t.pos.y)*1000 + (t.pos.x)*4 + FACING_NUM[t.facing]
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

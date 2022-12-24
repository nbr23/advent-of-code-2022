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
	pos       point
	facing    point
	facade_id int
}

var RIGHT = point{1, 0}
var LEFT = point{-1, 0}
var DOWN = point{0, 1}
var UP = point{0, -1}

var FACING_NUM = map[point]int{
	point{1, 0}:  0,
	point{0, 1}:  1,
	point{-1, 0}: 2,
	point{0, -1}: 3,
}

var FACING_CHAR = map[point]rune{
	point{1, 0}:  '>',
	point{0, 1}:  'v',
	point{-1, 0}: '<',
	point{0, -1}: '^',
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
		0, // doesn't matter in p1
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

type transition struct {
	facade_i  int
	transform func(p *point) (point, point)
}

var FACADES []map[point]transition

var FACADES_REAL = []map[point]transition{
	// 1
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{0, p.x}, RIGHT }},
		DOWN:  {5, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {3, func(p *point) (point, point) { return point{0, BOX_SIZE - 1 - p.y}, RIGHT }},
	},
	// 2
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {5, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.x}, LEFT }},
		RIGHT: {6, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - 1 - p.y}, LEFT }},
		LEFT:  {1, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 3
	map[point]transition{
		UP:    {5, func(p *point) (point, point) { return point{0, p.x}, RIGHT }},
		DOWN:  {4, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {6, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {1, func(p *point) (point, point) { return point{0, BOX_SIZE - p.y - 1}, RIGHT }},
	},
	// 4
	map[point]transition{
		UP:    {3, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {2, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {6, func(p *point) (point, point) { return point{p.y, BOX_SIZE - 1}, UP }},
		LEFT:  {1, func(p *point) (point, point) { return point{p.y, 0}, DOWN }},
	},
	// 5
	map[point]transition{
		UP:    {1, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {6, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{p.y, BOX_SIZE - 1}, UP }},
		LEFT:  {3, func(p *point) (point, point) { return point{p.y, 0}, DOWN }},
	},
	// 6
	map[point]transition{
		UP:    {5, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {4, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.x}, LEFT }},
		RIGHT: {2, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - 1 - p.y}, LEFT }},
		LEFT:  {3, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
}

var FACADES_TEST = []map[point]transition{
	// 1
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{BOX_SIZE - p.x - 1, 0}, DOWN }},
		DOWN:  {5, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - 1 - p.y}, LEFT }},
		LEFT:  {3, func(p *point) (point, point) { return point{p.y, 0}, DOWN }},
	},
	// 2
	map[point]transition{
		UP:    {5, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - p.x - 1}, LEFT }},
		DOWN:  {4, func(p *point) (point, point) { return point{0, p.x}, RIGHT }},
		RIGHT: {1, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - 1 - p.y}, LEFT }},
		LEFT:  {6, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 3
	map[point]transition{
		UP:    {1, func(p *point) (point, point) { return point{0, p.x}, RIGHT }},
		DOWN:  {6, func(p *point) (point, point) { return point{0, BOX_SIZE - p.x - 1}, RIGHT }},
		RIGHT: {5, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {4, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 4
	map[point]transition{
		UP:    {1, func(p *point) (point, point) { return point{BOX_SIZE - 1 - p.x, 0}, DOWN }},
		DOWN:  {6, func(p *point) (point, point) { return point{BOX_SIZE - 1 - p.x, BOX_SIZE - 1}, UP }},
		RIGHT: {3, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {2, func(p *point) (point, point) { return point{BOX_SIZE - 1 - p.y, BOX_SIZE - 1}, UP }},
	},
	// 5
	map[point]transition{
		UP:    {1, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {6, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{BOX_SIZE - 1 - p.y, 0}, DOWN }},
		LEFT:  {3, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 6
	map[point]transition{
		UP:    {5, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {4, func(p *point) (point, point) { return point{BOX_SIZE - 1 - p.x, BOX_SIZE - 1}, UP }},
		RIGHT: {2, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {3, func(p *point) (point, point) { return point{BOX_SIZE - p.y - 1, BOX_SIZE - 1}, UP }},
	},
}

var FACADES_DUMMY = []map[point]transition{
	// 1
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {5, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {3, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 2
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - p.x - 1}, LEFT }},
		DOWN:  {5, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.x}, LEFT }},
		RIGHT: {6, func(p *point) (point, point) { return point{0, BOX_SIZE - 1 - p.y}, LEFT }},
		LEFT:  {1, func(p *point) (point, point) { return point{BOX_SIZE - 1, p.y}, LEFT }},
	},
	// 3
	map[point]transition{
		UP:    {4, func(p *point) (point, point) { return point{0, p.x}, RIGHT }},
		DOWN:  {5, func(p *point) (point, point) { return point{0, BOX_SIZE - p.x - 1}, RIGHT }},
		RIGHT: {1, func(p *point) (point, point) { return point{0, p.y}, RIGHT }},
		LEFT:  {6, func(p *point) (point, point) { return point{0, BOX_SIZE - p.y - 1}, RIGHT }},
	},
	// 4
	map[point]transition{
		UP:    {6, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {1, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{BOX_SIZE - p.y - 1, 0}, DOWN }},
		LEFT:  {3, func(p *point) (point, point) { return point{p.y, 0}, DOWN }},
	},
	// 5
	map[point]transition{
		UP:    {1, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {6, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{p.y, BOX_SIZE - 1}, UP }},
		LEFT:  {3, func(p *point) (point, point) { return point{BOX_SIZE - p.y - 1, BOX_SIZE - 1}, UP }},
	},
	// 6
	map[point]transition{
		UP:    {5, func(p *point) (point, point) { return point{p.x, BOX_SIZE - 1}, UP }},
		DOWN:  {4, func(p *point) (point, point) { return point{p.x, 0}, DOWN }},
		RIGHT: {2, func(p *point) (point, point) { return point{BOX_SIZE - 1, BOX_SIZE - p.y - 1}, LEFT }},
		LEFT:  {3, func(p *point) (point, point) { return point{0, BOX_SIZE - p.y - 1}, RIGHT }},
	},
}

func setFacadePattern(box_size int) {
	if box_size == 4 { // we're in test mode
		FACADE_PATTERNS = [][]int{
			{0, 0, 1, 0},
			{4, 3, 5, 0},
			{0, 0, 6, 2},
		}
	} else { // REAL_MODE!
		FACADE_PATTERNS = [][]int{
			{0, 1, 2},
			{0, 5, 0},
			{3, 6, 0},
			{4, 0, 0},
		}
	}
}

func getFaceCoords(face int) point {
	for y := 0; y < len(FACADE_PATTERNS); y++ {
		for x := 0; x < len(FACADE_PATTERNS[y]); x++ {
			if FACADE_PATTERNS[y][x] == face {
				return point{x * BOX_SIZE, y * BOX_SIZE}
			}
		}
	}
	return point{}
}

func printFacades(facades map[int]map[point]bool, box_size int) {
	for fy := 0; fy < len(FACADE_PATTERNS); fy++ {
		for y := 0; y < box_size; y++ {
			for fx := 0; fx < len(FACADE_PATTERNS[0]); fx++ {
				for x := 0; x < box_size; x++ {
					printed := false
					if FACADE_PATTERNS[fy][fx] == 0 {
						fmt.Print(" ")
						continue
					}
					for _, t := range TRAIL {
						if x == t.coords.x && y == t.coords.y && FACADE_PATTERNS[fy][fx] == t.facade {
							fmt.Print(fmt.Sprintf("%c", t.icon))
							printed = true
							continue
						}
					}
					if printed {
						continue
					}
					if facades[FACADE_PATTERNS[fy][fx]][point{x, y}] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
			}
			fmt.Println()
		}
	}
}

func getFacades(monkeyMap map[point]bool, box_size int) map[int]map[point]bool {
	facades := make(map[int]map[point]bool)
	for fy := 0; fy < len(FACADE_PATTERNS); fy++ {
		for fx := 0; fx < len(FACADE_PATTERNS[0]); fx++ {
			if FACADE_PATTERNS[fy][fx] == 0 { // skip things that aren't facades
				continue
			}
			newface := make(map[point]bool)
			for x := 0; x < box_size; x++ {
				for y := 0; y < box_size; y++ {
					newface[point{x, y}] = monkeyMap[point{fx*box_size + x, fy*box_size + y}]
				}
			}
			facades[FACADE_PATTERNS[fy][fx]] = newface
		}
	}
	return facades
}

var FACADE_PATTERNS [][]int

type trail struct {
	icon   rune
	coords point
	facade int
}

var TRAIL = make([]trail, 0)

func moveTurtleCube(monkeyMap map[int]map[point]bool, t *turtle, steps int) {
	for i := 0; i < steps; i++ {
		// face changes
		if t.pos.x+t.facing.x >= BOX_SIZE || t.pos.x+t.facing.x < 0 || t.pos.y+t.facing.y >= BOX_SIZE || t.pos.y+t.facing.y < 0 {
			newf := FACADES[t.facade_id-1][t.facing].facade_i
			newp, newfacing := FACADES[t.facade_id-1][t.facing].transform(&point{t.pos.x, t.pos.y})
			v := monkeyMap[newf][newp]
			if v {
				return // we're done, we hit a wall
			}
			t.facade_id = newf
			t.pos.x = newp.x
			t.pos.y = newp.y
			t.facing = newfacing
		} else { // move within current face, all chill
			newp := point{t.pos.x + t.facing.x, t.pos.y + t.facing.y}
			v := monkeyMap[t.facade_id][newp]
			if v {
				return // we're done, we hit a wall
			}
			t.pos.x = newp.x
			t.pos.y = newp.y
		}
		TRAIL = append(TRAIL, trail{FACING_CHAR[t.facing], t.pos, t.facade_id})
	}
}

var BOX_SIZE = 0

func part2(input string) interface{} {
	parts := strings.Split(strings.Trim(input, "\n"), "\n\n")

	angles, steps := splitDirectionsChunks(parts[1])

	monkeyMap_s := strings.Split(parts[0], "\n")
	monkeyMap := loadMap(monkeyMap_s)

	t := turtle{
		point{0, 0},
		point{1, 0},
		1,
	}

	BOX_SIZE = utils.IntMax(getMaxX(monkeyMap_s), len(monkeyMap_s)) / 4
	if BOX_SIZE == 4 {
		FACADES = FACADES_TEST
	} else {
		FACADES = FACADES_REAL
	}

	setFacadePattern(BOX_SIZE)
	facades := getFacades(monkeyMap, BOX_SIZE)

	for i, step := range steps {
		moveTurtleCube(facades, &t, step)
		if i < len(angles) {
			rotate(&t.facing, angles[i])
		}
	}
	t.pos.y++
	t.pos.x++
	box := getFaceCoords(t.facade_id)
	return (t.pos.y+box.y)*1000 + (t.pos.x+box.x)*4 + FACING_NUM[t.facing]
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

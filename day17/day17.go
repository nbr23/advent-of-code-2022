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
	width     int
	height    int
	rockfun   RockFunc
	rockfun_i RockFun_i
	points    []*Point
	points_i  []int
}

// -
func Rock_1(coord Point, pts *[]*Point) {
	(*pts)[0].X, (*pts)[0].Y = coord.X, coord.Y
	(*pts)[1].X, (*pts)[1].Y = coord.X+1, coord.Y
	(*pts)[2].X, (*pts)[2].Y = coord.X+2, coord.Y
	(*pts)[3].X, (*pts)[3].Y = coord.X+3, coord.Y
}
func Rock_1_i(coord Point, pts *[]int) {
	(*pts)[0] = coord.Y*7 + coord.X
	(*pts)[1] = coord.Y*7 + coord.X + 1
	(*pts)[2] = coord.Y*7 + coord.X + 2
	(*pts)[3] = coord.Y*7 + coord.X + 3
}

var ROCK_1 = Rock{
	4,
	1,
	Rock_1,
	Rock_1_i,
	make([]*Point, 4),
	make([]int, 4),
}

// +
func Rock_2(coord Point, pts *[]*Point) {
	(*pts)[0].X, (*pts)[0].Y = coord.X+1, coord.Y
	(*pts)[1].X, (*pts)[1].Y = coord.X, coord.Y+1
	(*pts)[2].X, (*pts)[2].Y = coord.X+2, coord.Y+1
	(*pts)[3].X, (*pts)[3].Y = coord.X+1, coord.Y+2
}

func Rock_2_i(coord Point, pts *[]int) {
	(*pts)[0] = coord.Y*7 + coord.X + 1
	(*pts)[1] = (coord.Y+1)*7 + coord.X
	(*pts)[2] = (coord.Y+1)*7 + coord.X + 2
	(*pts)[3] = (coord.Y+2)*7 + coord.X + 1
}

var ROCK_2 = Rock{
	3,
	1,
	Rock_2,
	Rock_2_i,
	make([]*Point, 4),
	make([]int, 4),
}

// L backwards
func Rock_3(coord Point, pts *[]*Point) {
	(*pts)[0].X, (*pts)[0].Y = coord.X+2, coord.Y+2
	(*pts)[1].X, (*pts)[1].Y = coord.X+2, coord.Y+1
	(*pts)[2].X, (*pts)[2].Y = coord.X, coord.Y
	(*pts)[3].X, (*pts)[3].Y = coord.X+1, coord.Y
	(*pts)[4].X, (*pts)[4].Y = coord.X+2, coord.Y
}
func Rock_3_i(coord Point, pts *[]int) {
	(*pts)[0] = (coord.Y+2)*7 + coord.X + 2
	(*pts)[1] = (coord.Y+1)*7 + coord.X + 2
	(*pts)[2] = (coord.Y)*7 + coord.X
	(*pts)[3] = (coord.Y)*7 + coord.X + 1
	(*pts)[4] = (coord.Y)*7 + coord.X + 2
}

var ROCK_3 = Rock{
	3,
	1,
	Rock_3,
	Rock_3_i,
	make([]*Point, 5),
	make([]int, 5),
}

// l
func Rock_4(coord Point, pts *[]*Point) {
	(*pts)[0].X, (*pts)[0].Y = coord.X, coord.Y
	(*pts)[1].X, (*pts)[1].Y = coord.X, coord.Y+1
	(*pts)[2].X, (*pts)[2].Y = coord.X, coord.Y+2
	(*pts)[3].X, (*pts)[3].Y = coord.X, coord.Y+3
}
func Rock_4_i(coord Point, pts *[]int) {
	(*pts)[0] = (coord.Y)*7 + coord.X
	(*pts)[1] = (coord.Y+1)*7 + coord.X
	(*pts)[2] = (coord.Y+2)*7 + coord.X
	(*pts)[3] = (coord.Y+3)*7 + coord.X
}

var ROCK_4 = Rock{
	1,
	1,
	Rock_4,
	Rock_4_i,
	make([]*Point, 4),
	make([]int, 4),
}

// square
func Rock_5(coord Point, pts *[]*Point) {
	(*pts)[0].X, (*pts)[0].Y = coord.X, coord.Y
	(*pts)[1].X, (*pts)[1].Y = coord.X+1, coord.Y
	(*pts)[2].X, (*pts)[2].Y = coord.X, coord.Y+1
	(*pts)[3].X, (*pts)[3].Y = coord.X+1, coord.Y+1
}
func Rock_5_i(coord Point, pts *[]int) {
	(*pts)[0] = (coord.Y)*7 + coord.X
	(*pts)[1] = (coord.Y)*7 + coord.X + 1
	(*pts)[2] = (coord.Y+1)*7 + coord.X
	(*pts)[3] = (coord.Y+1)*7 + coord.X + 1
}

var ROCK_5 = Rock{
	2,
	1,
	Rock_5,
	Rock_5_i,
	make([]*Point, 4),
	make([]int, 4),
}

type RockFunc func(coord Point, pts *[]*Point)
type RockFun_i func(coord Point, pts *[]int)

var ROCKS = []Rock{
	ROCK_1,
	ROCK_2,
	ROCK_3,
	ROCK_4,
	ROCK_5,
}

func canGoTo(playground Playground, current_rock *Rock, current_rock_coords Point) bool {
	current_rock.rockfun_i(current_rock_coords, &current_rock.points_i)
	for _, coords := range current_rock.points_i {
		if _, ok := playground[coords]; ok {
			return false
		}
	}
	return true
}

func fillFloor(playground Playground) {
	for i := 0; i <= 7; i++ {
		playground[i] = true
	}
}

func printMap(playground Playground, highest int, current Rock, current_pos Point) {
	for y := highest; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if _, ok := playground[y*7+x]; ok {
				fmt.Print("#")
			} else {
				done := false
				current.rockfun(current_pos, &current.points)
				for _, pos := range current.points {
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

// func garbageCollect(playground Playground, highest int) {
// 	newfloor := highest
// 	for x := 0; x < 7; x++ {
// 		for y := newfloor; y > 0; y-- {
// 			if _, ok := playground[y*7+x]; ok {
// 				newfloor = utils.IntMin(newfloor, y)
// 				break
// 			}
// 		}
// 	}
// 	// fmt.Printf("Newfloor %d\n", newfloor)
// 	for pt := range playground { // y*7+x
// 		if pt < newfloor*7 {
// 			delete(playground, pt)
// 		}
// 	}
// }

type Playground map[int]bool

func findPattern(heights []int, cycle int) (int, int, int) {

	for p_len := cycle; p_len < cycle*10; p_len++ {
		for p_start := cycle; p_start+2*p_len < len(heights); p_start++ {
			good := false
			diff := heights[p_start+p_len] - heights[p_start]

			for i := p_start; i < p_start+p_len; i++ {
				currdiff := heights[i+p_len] - heights[i]
				if currdiff != diff {
					good = false
					break
				}
				good = true
			}
			if good {
				return p_start, p_len, diff
			}
		}
	}
	return -1, -1, -1
}

func dropRocks(input string, MAX int) int {
	playground := make(Playground)
	heights := make([]int, 0)
	fillFloor(playground)
	rockindex := 0
	highest := 0
	var current_rock *Rock
	current_rock = nil
	var current_rock_coord Point

	rock_count := MAX + 1
	for i := 0; ; i++ {
		if current_rock == nil {
			current_rock = &ROCKS[rockindex]
			current_rock_coord = Point{2, highest + current_rock.height + 3}
			rockindex = (rockindex + 1) % len(ROCKS)
			rock_count--

			heights = append(heights, highest)
			if rock_count == 0 {
				return highest
			}
			if len(input)*5*10 == MAX-rock_count {
				pattern_start, pattern_len, diff := findPattern(heights, len(input)*5)
				periodsCount := (MAX - pattern_start) / pattern_len
				periodsRest := (MAX - pattern_start) % pattern_len
				return heights[pattern_start] + heights[periodsRest+pattern_start] - heights[pattern_start] + diff*periodsCount

			}
		}

		c := input[i%len(input)]

		// Lateral move
		if c == '>' {
			if current_rock_coord.X+1+current_rock.width-1 < 7 && canGoTo(playground, current_rock, Point{current_rock_coord.X + 1, current_rock_coord.Y}) {
				current_rock_coord.X++
			} else {
			}
		} else {
			if current_rock_coord.X-1 >= 0 && canGoTo(playground, current_rock, Point{current_rock_coord.X - 1, current_rock_coord.Y}) {
				current_rock_coord.X--
			} else {
			}
		}

		// down
		if canGoTo(playground, current_rock, Point{current_rock_coord.X, current_rock_coord.Y - 1}) {
			current_rock_coord.Y--
		} else {
			current_rock.rockfun_i(current_rock_coord, &current_rock.points_i)
			for _, rock := range current_rock.points_i {
				highest = utils.IntMax(highest, rock/7)
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

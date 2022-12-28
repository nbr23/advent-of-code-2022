package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 24

type point struct {
	x int
	y int
	z int // z is time
}

type blizzardMap struct {
	bmap map[point]bool
	maxx int
	maxy int
}

type blizzard struct {
	position  point
	direction *point
}

var RIGHT = point{1, 0, 1}
var LEFT = point{-1, 0, 1}
var DOWN = point{0, 1, 1}
var UP = point{0, -1, 1}
var HERE = point{0, 0, 1}

var MOVES = []point{
	RIGHT,
	LEFT,
	DOWN,
	UP,
	HERE,
}

func getDirectionIcon(p point) string {
	if p.x == 1 {
		return ">"
	} else if p.x == -1 {
		return "<"
	} else if p.y == 1 {
		return "v"
	} else if p.y == -1 {
		return "^"
	}
	panic("Wrong way")
}

func getDirection(c rune) *point {
	if c == '^' {
		return &UP
	} else if c == '<' {
		return &LEFT
	} else if c == '>' {
		return &RIGHT
	} else if c == 'v' {
		return &DOWN
	}
	panic("Wrong way")
}

func mod(i, n int) int {
	return ((i % n) + n) % n
}

func parseMap(input string) *blizzardMap {
	var blizmap blizzardMap
	blizmap.bmap = make(map[point]bool)
	ilist := inputs.InputToStrList(input)
	blizmap.maxy = len(ilist)
	blizmap.maxx = len(ilist[0])
	for y := range ilist {
		for x, c := range ilist[y] {
			if c == '>' || c == 'v' || c == '<' || c == '^' {
				b := blizzard{point{x, y, 0}, getDirection(c)}
				for z := 0; z <= (blizmap.maxx-2)*(blizmap.maxy-2); z++ {
					blizmap.bmap[getBlizzardPosition(&blizmap, &b, z)] = true
				}
			}
		}
	}
	return &blizmap
}

// get the blizzard's position at the current minute
func getBlizzardPosition(bmap *blizzardMap, b *blizzard, minute int) point {
	return point{
		mod(b.position.x+(b.direction.x*minute)-1, bmap.maxx-2) + 1,
		mod(b.position.y+(b.direction.y*minute)-1, bmap.maxy-2) + 1,
		minute,
	}
}

func pathFinder(blizmap *blizzardMap, start, end point, minutes int) int {
	current_positions := mapset.NewSet[point]()
	current_positions.Add(start)
	cost := minutes

	for {
		new_current_positions := mapset.NewSet[point]()
		for p := range current_positions.Iter() {
			for _, m := range MOVES {
				newpoint := point{p.x + m.x, p.y + m.y, (cost + 1) % ((blizmap.maxx - 2) * (blizmap.maxy - 2))} // fixme or modulo here ?

				if newpoint.x == end.x && newpoint.y == end.y {
					return cost + 1
				}
				if !(newpoint.x == start.x && newpoint.y == start.y) && (newpoint.x < 1 || newpoint.x >= blizmap.maxx-1 || newpoint.y < 1 || newpoint.y >= blizmap.maxy-1) {
					continue
				}

				if v, ok := blizmap.bmap[newpoint]; ok || v {
					continue
				}

				new_current_positions.Add(newpoint)
			}
		}
		current_positions.Clear()
		current_positions = new_current_positions
		cost++
	}
}

var P1 int
var bmap *blizzardMap

func part1(input string) interface{} {
	bmap = parseMap(input)
	P1 = pathFinder(bmap, point{1, 0, 0}, point{bmap.maxx - 2, bmap.maxy - 1, 0}, 0)
	return P1
}

func part2(input string) interface{} {
	back := pathFinder(bmap, point{bmap.maxx - 2, bmap.maxy - 1, 0}, point{1, 0, 0}, P1)
	return pathFinder(bmap, point{1, 0, 0}, point{bmap.maxx - 2, bmap.maxy - 1, 0}, back)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

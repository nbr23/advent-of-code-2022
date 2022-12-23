package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 23

type point struct {
	x int
	y int
}

type elfMap struct {
	e_map  map[point]bool
	e_list []*point
	minx   int
	miny   int
	maxx   int
	maxy   int
}

func hasNeighbors(m *elfMap, e *point) bool {
	for x := e.x - 1; x <= e.x+1; x++ {
		for y := e.y - 1; y <= e.y+1; y++ {
			if e.x == x && e.y == y {
				continue
			}
			v, ok := m.e_map[point{x, y}]
			if ok && v {
				return true
			}
		}
	}
	return false
}

func printMap(m *elfMap) {
	for y := m.miny; y <= m.maxy; y++ {
		for x := m.minx; x <= m.maxx; x++ {
			if m.e_map[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseElfMap(input string) *elfMap {
	slist := inputs.InputToStrList(input)
	emap := elfMap{make(map[point]bool), make([]*point, 0), 0, 0, len(slist[0]), len(slist)}
	for y, s := range slist {
		for x, c := range s {
			emap.e_map[point{x, y}] = c == '#'
			if c == '#' {
				emap.e_list = append(emap.e_list, &point{x, y})
			}
		}
	}
	return &emap
}

var CONSIDERED_POSITIONS = [][]point{
	// NE, N, NW
	[]point{
		{-1, -1}, {0, -1}, {1, -1},
	},

	// SE, S, SW
	[]point{
		{-1, 1}, {0, 1}, {1, 1},
	},
	// W, NW, SE
	[]point{
		{-1, -1}, {-1, 0}, {-1, 1},
	},

	// E, NE, SE
	[]point{
		{1, -1}, {1, 0}, {1, 1},
	},
}

func getMinMax(m *elfMap) {
	m.minx, m.maxx, m.miny, m.maxy = m.e_list[0].x, m.e_list[0].x, m.e_list[0].y, m.e_list[0].y
	for _, elf := range m.e_list {
		m.minx = utils.IntMin(m.minx, elf.x)
		m.maxx = utils.IntMax(m.maxx, elf.x)
		m.miny = utils.IntMin(m.miny, elf.y)
		m.maxy = utils.IntMax(m.maxy, elf.y)
	}
}

var P1 = 0
var P2 = 0

func compute(input string) {
	emap := parseElfMap(input)
	getMinMax(emap)

	round := 0

	for ; ; round++ {
		if round == 10 {
			P1 = (emap.maxx-emap.minx+1)*(emap.maxy-emap.miny+1) - len(emap.e_list)
		}

		moves := make(map[point][]*point)

		for _, elf := range emap.e_list {
			if hasNeighbors(emap, elf) {
				for i := 0; i < len(CONSIDERED_POSITIONS); i++ {
					poss := CONSIDERED_POSITIONS[(round+i)%len(CONSIDERED_POSITIONS)]
					canmove := true
					for _, p := range poss {
						v, ok := emap.e_map[point{elf.x + p.x, elf.y + p.y}]
						if ok && v {
							canmove = false
							break
						}
					}
					// we found out we can move in that direction of poss[1], we
					if canmove {
						newpos := point{elf.x + poss[1].x, elf.y + poss[1].y}
						_, ok := moves[newpos]
						if !ok {
							moves[newpos] = make([]*point, 0)
						}
						moves[newpos] = append(moves[newpos], elf)
						break
					}
				}
			}
		}

		// now we actually perform the moves
		moves_count := 0
		for p := range moves {
			if len(moves[p]) != 1 {
				continue
			}
			emap.e_map[*moves[p][0]] = false
			moves[p][0].x = p.x
			moves[p][0].y = p.y
			emap.e_map[*moves[p][0]] = true
			moves_count++
		}

		if moves_count == 0 {
			P2 = round + 1
			return
		}

		getMinMax(emap)
	}
}

func part1(input string) interface{} {
	compute(input)
	return P1
}

func part2(input string) interface{} {
	return P2
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

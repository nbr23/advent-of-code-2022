package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"math"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 12

type Point struct {
	x int
	y int
}

func loadInput(input string) (map[Point]int, Point, Point) {
	input = strings.Trim(input, "\n")
	split := strings.Split(input, "\n")
	S := Point{}
	E := Point{}
	matrice := make(map[Point]int)
	for i, s := range split {
		for j, c := range s {
			if c == 'S' {
				S = Point{j, i}
				matrice[S] = 0
			} else if c == 'E' {
				E = Point{j, i}
				matrice[E] = 25
			} else {
				matrice[Point{j, i}] = int(c - 'a')
			}
		}
	}
	return matrice, S, E
}

var LOOKAROUND = []Point{
	Point{1, 0},
	Point{0, 1},
	Point{-1, 0},
	Point{0, -1},
}

func computeCosts(matrice map[Point]int, costs map[Point]int, cursor Point, end Point, cost int) {

	if cursor.x == end.x && cursor.y == end.y {
		return
	}

	height := matrice[cursor]

	for _, p := range LOOKAROUND {
		newcursor := Point{cursor.x + p.x, cursor.y + p.y}
		if h, ok := matrice[newcursor]; ok {
			// valid point
			if h <= height+1 {
				// check if already have a cost for it
				if ecost, ok := costs[newcursor]; ok {
					if cost+1 < ecost {
						costs[newcursor] = cost + 1
						computeCosts(matrice, costs, newcursor, end, cost+1)
					}
				} else {
					costs[newcursor] = cost + 1
					computeCosts(matrice, costs, newcursor, end, cost+1)
				}
			}
		}
	}
}

func part1(input string) interface{} {
	costs := make(map[Point]int)
	matrice, Start, End := loadInput(input)
	computeCosts(matrice, costs, Start, End, 0)
	return costs[End]
}

func computeCostsp2(matrice map[Point]int, costs map[Point]int, cursor Point, cost int) {
	height := matrice[cursor]

	for _, p := range LOOKAROUND {
		newcursor := Point{cursor.x + p.x, cursor.y + p.y}
		if h, ok := matrice[newcursor]; ok {
			// valid point
			if height <= h+1 {
				// check if already have a cost for it
				if ecost, ok := costs[newcursor]; ok {
					if cost+1 < ecost {
						costs[newcursor] = cost + 1
						computeCostsp2(matrice, costs, newcursor, cost+1)
					}
				} else {
					costs[newcursor] = cost + 1
					computeCostsp2(matrice, costs, newcursor, cost+1)
				}
			}
		}
	}
}

func findNearestA(matrice, costs map[Point]int) int {
	cost := math.MaxInt
	for p, v := range matrice {
		if v == 0 && costs[p] < cost && costs[p] != 0 {
			cost = costs[p]
		}
	}
	return cost
}

func part2(input string) interface{} {
	costs := make(map[Point]int)
	matrice, _, End := loadInput(input)
	computeCostsp2(matrice, costs, End, 0)
	return findNearestA(matrice, costs)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

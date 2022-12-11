package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 9

type Point struct {
	x, y int
}

func getVector(dir string) Point {
	return map[string]Point{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}[dir]
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func part1(input string) interface{} {
	var T Point
	var H Point
	visited := make(map[Point]bool)
	visited[T] = true

	for _, line := range inputs.InputToStrList(input) {
		instr := strings.Split(line, " ")
		moves := inputs.ParseDecInt(instr[1])
		for i := 0; i < moves; i++ {
			v := getVector(instr[0])
			H.x += v.x
			H.y += v.y
			// Move tail
			diff := Point{abs(T.x - H.x), abs(T.y - H.y)}
			if diff.x <= 1 && diff.y <= 1 {
				// no action
				continue
			}
			// aligned, so just increment as needed
			if diff.x == 0 && H.y > T.y {
				T.y++
			} else if diff.x == 0 && H.y < T.y {
				T.y--
			} else if diff.y == 0 && H.x > T.x {
				T.x++
			} else if diff.y == 0 && H.x < T.x {
				T.x--
			} else {
				// Diagonals
				if H.x > T.x {
					T.x++
				} else {
					T.x--
				}
				if H.y > T.y {
					T.y++
				} else {
					T.y--
				}
			}
			visited[T] = true
		}
	}
	return len(visited)
}

func part2(input string) interface{} {
	Tails := make([]Point, 9)
	var Head Point
	visited := make(map[Point]bool)
	visited[Tails[8]] = true

	for _, line := range inputs.InputToStrList(input) {
		instr := strings.Split(line, " ")
		moves := inputs.ParseDecInt(instr[1])
		for i := 0; i < moves; i++ {
			v := getVector(instr[0])
			Head.x += v.x
			Head.y += v.y
			// Move tail
			for ti := range Tails {
				var H *Point
				var T *Point
				if ti == 0 {
					H = &Head
				} else {
					H = &Tails[ti-1]
				}
				T = &Tails[ti]

				diff := Point{abs(T.x - H.x), abs(T.y - H.y)}
				if diff.x <= 1 && diff.y <= 1 {
					// no action
					continue
				}
				// aligned, so just increment as needed
				if diff.x == 0 && H.y > T.y {
					T.y++
				} else if diff.x == 0 && H.y < T.y {
					T.y--
				} else if diff.y == 0 && H.x > T.x {
					T.x++
				} else if diff.y == 0 && H.x < T.x {
					T.x--
				} else {
					// Diagonals
					if H.x > T.x {
						T.x++
					} else {
						T.x--
					}
					if H.y > T.y {
						T.y++
					} else {
						T.y--
					}
				}
				if ti == 8 {
					visited[Tails[ti]] = true
				}
			}
		}
	}
	return len(visited)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

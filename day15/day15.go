package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 15

type Point struct {
	x, y int
}

type Map struct {
	maxX, maxY, minX, minY int
	sensors                map[Point]int
	beacons                map[Point]bool
}

func parseBeacons(input string) Map {
	var beaconsMap Map
	beaconsMap.sensors = make(map[Point]int)
	beaconsMap.beacons = make(map[Point]bool)
	for _, line := range inputs.InputToStrList(input) {
		var p1, p2 Point
		parts := strings.Split(line, " ")
		p1.x = inputs.ParseDecInt(strings.Trim(strings.Split(parts[2], "=")[1], ","))
		p1.y = inputs.ParseDecInt(strings.Trim(strings.Split(parts[3], "=")[1], ":"))

		p2.x = inputs.ParseDecInt(strings.Trim(strings.Split(parts[8], "=")[1], ","))
		p2.y = inputs.ParseDecInt(strings.Split(parts[9], "=")[1])

		distance := manhattanDist(p1, p2)

		beaconsMap.maxX = utils.IntMax(beaconsMap.maxX, p1.x+distance)
		beaconsMap.maxY = utils.IntMax(beaconsMap.maxY, p1.y+distance)

		beaconsMap.minX = utils.IntMin(beaconsMap.minX, p1.x-distance)
		beaconsMap.minY = utils.IntMin(beaconsMap.minY, p1.y-distance)

		beaconsMap.sensors[p1] = distance
		beaconsMap.beacons[p2] = true
		drawCircle(p1, distance)
	}
	return beaconsMap
}

func manhattanDist(p1, p2 Point) int {
	return utils.IntAbs(p1.x-p2.x) + utils.IntAbs(p1.y-p2.y)
}

func part1(input string) interface{} {
	Y := 2000000
	beaconsMap = parseBeacons(input)
	score := 0
	scan := Point{beaconsMap.minX, Y}

	for ; scan.x <= beaconsMap.maxX; scan.x++ {
		if _, ok := beaconsMap.beacons[scan]; ok {
			continue
		}
		if _, ok := beaconsMap.sensors[scan]; ok {
			continue
		}
		for p, d := range beaconsMap.sensors {
			if manhattanDist(scan, p) <= d {
				score++
				break
			}
		}
	}
	return score
}

var fullMap = make(map[Point]bool)

var ranges = make([][]*Point, MAX*2)

func (pts *Point) String() string {
	return fmt.Sprintf("%v", *pts)
}

func addRange(y, min, max int) {
	if ranges[y] == nil {
		ranges[y] = []*Point{{min, max}}
		return
	}

	done := false
	for _, p := range ranges[y] {
		if min <= p.y && min >= p.x {
			p.y = utils.IntMax(max, p.y)
			done = true
		}
		if max <= p.y && max >= p.y {
			p.x = utils.IntMin(p.x, min)
			done = true
		}
		if done {
			return
		}
	}
	ranges[y] = append(ranges[y], &Point{min, max})
}

func drawCircle(p Point, d int) {
	for h := 0; h <= d; h++ {
		if p.y+d-h >= 0 {
			addRange(p.y+d-h, p.x-h, p.x+h)
		}
		if p.y-d+h >= 0 {
			addRange(p.y-d+h, p.x-h, p.x+h)
		}
	}
}

var beaconsMap Map
var MAX = 4000000

func printMap() {
	for y := 0; y <= MAX; y++ {
		for x := 0; x <= MAX; x++ {
			found := false
			for _, r := range ranges[y] {
				if x <= r.y && x >= r.x {
					found = true
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findEmpty() int {
	score := 0
	for y := 0; y <= MAX; y++ {
		for x := 0; x <= MAX; {
			found := false
			for _, r := range ranges[y] {
				if x <= r.y && x >= r.x {
					found = true
					x = r.y
					break
				}
			}
			if !found {
				score = x*4000000 + y
			}
			x++
		}
	}
	return score
}

func part2(input string) interface{} {
	return findEmpty()
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

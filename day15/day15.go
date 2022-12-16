package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 15

//Sensor at x=2, y=18: closest beacon is at x=-2, y=15

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
	}
	return beaconsMap
}

func manhattanDist(p1, p2 Point) int {
	return utils.IntAbs(p1.x-p2.x) + utils.IntAbs(p1.y-p2.y)
}

func part1(input string) interface{} {
	Y := 2000000
	beaconsMap := parseBeacons(input)
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

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

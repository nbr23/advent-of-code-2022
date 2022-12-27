package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"regexp"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 19

type blueprint struct {
	ore_robot_ores       int
	clay_robot_ore       int
	obsidian_robot_ore   int
	obsidian_robot_clay  int
	geode_robot_ore      int
	geode_robot_obsidian int
}

func parseBlueprints(input string) []blueprint {
	lines := inputs.InputToStrList(input)
	blueprints := make([]blueprint, len(lines))
	re := regexp.MustCompile(`^Blueprint [0-9]+: Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.$`)
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		blueprints[i] = blueprint{
			inputs.ParseDecInt(parts[1]),
			inputs.ParseDecInt(parts[2]),
			inputs.ParseDecInt(parts[3]),
			inputs.ParseDecInt(parts[4]),
			inputs.ParseDecInt(parts[5]),
			inputs.ParseDecInt(parts[6]),
		}
	}
	return blueprints
}

func part1(input string) interface{} {

	return parseBlueprints(input)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"regexp"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 19

type ores struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type robots ores
type blueprint ores

type State struct {
	Robots robots
	Ores   ores
}

func parseBlueprints(input string) [][]blueprint {
	lines := inputs.InputToStrList(input)
	blueprints := make([][]blueprint, len(lines))
	re := regexp.MustCompile(`^Blueprint [0-9]+: Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.$`)
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		blueprints[i] = make([]blueprint, 4)
		blueprints[i][0] = blueprint{
			ore:      inputs.ParseDecInt(parts[1]),
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		blueprints[i][1] = blueprint{
			ore:      inputs.ParseDecInt(parts[2]),
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		blueprints[i][2] = blueprint{
			ore:      inputs.ParseDecInt(parts[3]),
			clay:     inputs.ParseDecInt(parts[4]),
			obsidian: 0,
			geode:    0,
		}
		blueprints[i][3] = blueprint{
			ore:      inputs.ParseDecInt(parts[5]),
			clay:     0,
			obsidian: inputs.ParseDecInt(parts[6]),
			geode:    0,
		}
	}
	return blueprints
}

func hasResources(robot int, bp []blueprint, o *ores) bool {
	return bp[robot].ore <= o.ore &&
		bp[robot].clay <= o.clay &&
		bp[robot].obsidian <= o.obsidian &&
		bp[robot].geode <= o.geode
}

func buildRobot(robot int, bp []blueprint, state *State) State {
	r := robots{
		state.Robots.ore,
		state.Robots.clay,
		state.Robots.obsidian,
		state.Robots.geode,
	}
	newstate := State{
		r,
		ores{
			state.Ores.ore - bp[robot].ore,
			state.Ores.clay - bp[robot].clay,
			state.Ores.obsidian - bp[robot].obsidian,
			state.Ores.geode - bp[robot].geode,
		},
	}
	collect(&newstate)
	if robot == 0 {
		newstate.Robots.ore++
	}
	if robot == 1 {
		newstate.Robots.clay++
	}
	if robot == 2 {
		newstate.Robots.obsidian++
	}
	if robot == 3 {
		newstate.Robots.geode++
	}
	return newstate
}

func collect(s *State) {
	s.Ores.ore += s.Robots.ore
	s.Ores.clay += s.Robots.clay
	s.Ores.obsidian += s.Robots.obsidian
	s.Ores.geode += s.Robots.geode
}

var max = 0

func processMinute(bp int, bps [][]blueprint, minute int, state *State, seen map[State]bool, maxMinutes int) {
	if minute == maxMinutes-1 {
		if state.Ores.geode+state.Robots.geode < max {
			return
		}
	}
	if minute == maxMinutes {
		if state.Ores.geode > max {
			max = state.Ores.geode
		}
		return
	}
	if _, ok := seen[*state]; ok {
		return
	}
	seen[*state] = true

	// attempt to build robots
	for r := 0; r < 4; r++ {
		if hasResources(r, bps[bp], &state.Ores) {
			newstate := buildRobot(r, bps[bp], state)
			processMinute(bp, bps, minute+1, &newstate, seen, maxMinutes)
		}
	}
	collect(state)
	processMinute(bp, bps, minute+1, state, seen, maxMinutes)
}

func part1(input string) interface{} {
	bps := parseBlueprints(input)
	res := 0

	for bp, _ := range bps {
		max = 0
		seen := make(map[State]bool)
		processMinute(bp, bps, 0, &State{
			robots{
				1, 0, 0, 0,
			},
			ores{
				0, 0, 0, 0,
			},
		}, seen, 24)
		res += max * (bp + 1)
	}

	return res
}

func part2(input string) interface{} {
	bps := parseBlueprints(input)
	res := 1

	for bp := 0; bp < 3 && bp < len(bps); bp++ {
		max = 0
		seen := make(map[State]bool)
		processMinute(bp, bps, 0, &State{
			robots{
				1, 0, 0, 0,
			},
			ores{
				0, 0, 0, 0,
			},
		}, seen, 32)
		res *= max
	}

	return res
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

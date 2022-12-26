package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 16

type Valve struct {
	id          string
	flowrate    int
	open        bool
	leadsTo     []string
	leadsToWent []bool
}

func (v Valve) String() string {
	return fmt.Sprintf("{Rate=%d open=%v leadsTo=%s}", v.flowrate, v.open, v.leadsTo)
}

// Valve DF has flow rate=0; tunnels lead to valves ON, IC
func parseValves(input string) map[string]*Valve {
	valves := make(map[string]*Valve)
	for _, line := range inputs.InputToStrList(input) {
		parts := strings.Split(line, " ")

		valvespart := strings.Split(line, "valves ")
		var leadsTo []string
		if len(valvespart) > 1 {
			leadsTo = strings.Split(strings.Split(line, "valves ")[1], ", ")
		} else {
			leadsTo = []string{strings.Split(line, "valve ")[1]}
		}
		v := Valve{
			parts[1],
			inputs.ParseDecInt(strings.Trim(strings.Split(parts[4], "=")[1], ";")),
			false,
			leadsTo,
			make([]bool, len(leadsTo)),
		}

		valves[parts[1]] = &v
	}
	return valves
}

var Valves map[string]*Valve

type Tunnel struct {
	source string
	target string
}

type tunnelPath struct {
	source  *tunnelPath
	current string
	// targets  []string
	distance int
}

type path struct {
	from string
	to   string
}

var shortestPaths = make(map[path][]string)

func shortestPath(from, to string) []string {
	if v, ok := shortestPaths[path{from, to}]; ok {
		return v
	}
	visited := make(map[string]bool)

	tocheck := []*tunnelPath{{nil, from, 0}}
	tocheck_next := make([]*tunnelPath, 0)
	pathsSoFar := make([]*tunnelPath, 0)

	for len(visited) != len(Valves) {
		didsomething := false
		for _, check := range tocheck {
			if _, ok := visited[check.current]; ok {
				continue
			}
			visited[check.current] = true
			didsomething = true
			if check.current == to {
				pathsSoFar = append(pathsSoFar, check)
			} else {
				// we add its leadsto to tochecknext
				for _, lt := range Valves[check.current].leadsTo {
					tocheck_next = append(tocheck_next, &tunnelPath{check, lt, check.distance + 1})
				}
			}

		}
		if !didsomething {
			break
		}

		tocheck = tocheck_next
		tocheck_next = make([]*tunnelPath, 0)
	}

	best := pathsSoFar[0]

	for _, p := range pathsSoFar {
		if p.distance < best.distance {
			best = p
		}
	}

	res := make([]string, 0)
	for best != nil {
		res = append(res, best.current)
		best = best.source
	}
	shortestPaths[path{from, to}] = res
	return res
}

func getWorthiness(destFlow int, pathlength int, minutesleft int) int {
	return (minutesleft - (pathlength)) * destFlow
}

type position struct {
	minutes       int
	current_valve string
	flowrate      int
	open_valves   []string
}

func isIn(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}

func getZeroValves() []string {
	res := make([]string, 0)
	for v := range Valves {
		if Valves[v].flowrate == 0 {
			res = append(res, v)
		}
	}
	return res
}

func part1(input string) interface{} {
	Valves = parseValves(input)

	points := 0

	final_positions := make([]position, 0)

	zerovalves := getZeroValves()
	fmt.Println("zerovalves", zerovalves)

	active_positions := []position{
		{
			30,
			"AA",
			0,
			zerovalves,
		},
	}

	for {
		// we are done
		if len(active_positions) == 0 {
			break
		}

		new_active_pos := make([]position, 0)

		// iterate over current active positions
		for _, ap := range active_positions {
			// no time left, we are done
			if ap.minutes < 0 {
				// we add this to final_positions
				final_positions = append(final_positions, ap)
				continue
			}
			didsomething := false
			// otherwise, we look at all the unopened valves
			for v := range Valves {
				if isIn(ap.open_valves, v) {
					continue
				}

				// get the path between current and that one
				p := shortestPath(ap.current_valve, v)
				// fmt.Println(p)

				//get the worthiness of this path
				w := getWorthiness(Valves[v].flowrate, len(p), ap.minutes)
				if w < 0 {
					continue
				}
				didsomething = true

				newopenvalves := make([]string, len(ap.open_valves)+1)

				copy(newopenvalves, ap.open_valves)
				newopenvalves[len(ap.open_valves)] = v

				// and we add this to our "new_active_POS" list
				new_active_pos = append(new_active_pos, position{
					ap.minutes - len(p),
					v,
					ap.flowrate + w,
					newopenvalves,
				})
			}
			if !didsomething {
				final_positions = append(final_positions, ap)
			}
		}

		active_positions = new_active_pos
	}

	points = 0
	for _, fp := range final_positions {
		if fp.flowrate > points {
			points = fp.flowrate
		}
	}
	return points
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

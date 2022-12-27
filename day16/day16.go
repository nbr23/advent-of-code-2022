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
	id          int
	flowrate    int
	leadsTo     []int
	leadsToWent []bool
}

func valveID(name string) int {
	return (int(name[0])-int('A'))*100 + int(name[1]) - int('A')
}

func IDtoValve(v int) string {
	return fmt.Sprintf("%c%c", rune((v/100)+'A'), rune(v%100+'A'))
}

func IDstoValves(valves []int) []string {
	res := make([]string, len(valves))
	for i, v := range valves {
		res[i] = IDtoValve(v)
	}
	return res
}

func (v Valve) String() string {
	return fmt.Sprintf("{Rate=%d leadsTo=%v}", v.flowrate, v.leadsTo)
}

// Valve DF has flow rate=0; tunnels lead to valves ON, IC
func parseValves(input string) map[int]*Valve {
	valves := make(map[int]*Valve)
	for _, line := range inputs.InputToStrList(input) {
		parts := strings.Split(line, " ")

		valvespart := strings.Split(line, "valves ")
		var leadsTo = make([]int, 0)
		if len(valvespart) > 1 {
			for _, lt := range strings.Split(strings.Split(line, "valves ")[1], ", ") {
				leadsTo = append(leadsTo, valveID(lt))
			}
		} else {
			leadsTo = []int{valveID(strings.Split(line, "valve ")[1])}
		}
		v := Valve{
			valveID(parts[1]),
			inputs.ParseDecInt(strings.Trim(strings.Split(parts[4], "=")[1], ";")),
			leadsTo,
			make([]bool, len(leadsTo)),
		}

		valves[valveID(parts[1])] = &v
	}
	return valves
}

var Valves map[int]*Valve

type Tunnel struct {
	source int
	target int
}

type tunnelPath struct {
	source   *tunnelPath
	current  int
	distance int
}

type path struct {
	from int
	to   int
}

var shortestPaths = make(map[path]int)

func shortestPath(from, to int) int {
	if v, ok := shortestPaths[path{from, to}]; ok {
		return v
	}
	if v, ok := shortestPaths[path{to, from}]; ok {
		return v
	}
	visited := make(map[int]bool)

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

	res := make([]int, 0)
	for best != nil {
		res = append(res, best.current)
		best = best.source
	}
	shortestPaths[path{from, to}] = len(res)
	return len(res)
}

func getWorthiness(destFlow int, pathlength int, minutesleft int) int {
	return (minutesleft - (pathlength)) * destFlow
}

type position struct {
	minutes       int
	current_valve int
	flowrate      int
	open_valves   []int
}
type pair struct {
	s1 int
	s2 int
}

func isIn(l []int, s int) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}

func splitZeroValves() ([]int, []int) {
	zero := make([]int, 0)
	nonzero := make([]int, 0)
	for v := range Valves {
		if Valves[v].flowrate == 0 {
			zero = append(zero, v)
		} else {
			nonzero = append(nonzero, v)
		}
	}
	return zero, nonzero
}

func part1(input string) interface{} {
	Valves = parseValves(input)

	points := 0

	final_positions := make([]position, 0)

	zerovalves, nonzerovalves := splitZeroValves()

	active_positions := []position{
		{
			30,
			valveID("AA"),
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
			for _, v := range nonzerovalves {
				if isIn(ap.open_valves, v) {
					continue
				}

				// get the path between current and that one
				p := shortestPath(ap.current_valve, v)
				// fmt.Println(p)

				//get the worthiness of this path
				w := getWorthiness(Valves[v].flowrate, p, ap.minutes)
				if w < 0 {
					continue
				}
				didsomething = true

				newopenvalves := make([]int, len(ap.open_valves)+1)

				copy(newopenvalves, ap.open_valves)
				newopenvalves[len(ap.open_valves)] = v

				// and we add this to our "new_active_POS" list
				new_active_pos = append(new_active_pos, position{
					ap.minutes - p,
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
	Valves = parseValves(input)

	final_positions := make([]position, 0)

	zerovalves, nonzerovalves := splitZeroValves()

	active_positions := []position{
		{
			26,
			valveID("AA"),
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
			for _, v := range nonzerovalves {
				if isIn(ap.open_valves, v) {
					continue
				}

				// get the path between current and that one
				p := shortestPath(ap.current_valve, v)
				// fmt.Println(p)

				//get the worthiness of this path
				w := getWorthiness(Valves[v].flowrate, p, ap.minutes)
				if w < 0 {
					continue
				}
				didsomething = true

				newopenvalves := make([]int, len(ap.open_valves)+1)

				copy(newopenvalves, ap.open_valves)
				newopenvalves[len(ap.open_valves)] = v

				// and we add this to our "new_active_POS" list
				new_active_pos = append(new_active_pos, position{
					ap.minutes - p,
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

	humanpoints := 0
	var opened []int
	for _, fp := range final_positions {
		if fp.flowrate > humanpoints {
			humanpoints = fp.flowrate
			opened = make([]int, len(fp.open_valves))
			copy(opened, fp.open_valves)
		}
	}
	fmt.Println("opened", IDstoValves(opened))

	// starting elephant
	active_positions = []position{
		{
			26,
			valveID("AA"),
			0,
			opened,
		},
	}
	final_positions = make([]position, 0)

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
			for _, v := range nonzerovalves {
				if isIn(ap.open_valves, v) {
					continue
				}

				// get the path between current and that one
				p := shortestPath(ap.current_valve, v)
				// fmt.Println(p)

				//get the worthiness of this path
				w := getWorthiness(Valves[v].flowrate, p, ap.minutes)
				if w < 0 {
					continue
				}
				didsomething = true

				newopenvalves := make([]int, len(ap.open_valves)+1)

				copy(newopenvalves, ap.open_valves)
				newopenvalves[len(ap.open_valves)] = v

				// and we add this to our "new_active_POS" list
				new_active_pos = append(new_active_pos, position{
					ap.minutes - p,
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

	elephantpoints := 0
	for _, fp := range final_positions {
		if fp.flowrate > elephantpoints {
			elephantpoints = fp.flowrate
			opened = make([]int, len(fp.open_valves))
			copy(opened, fp.open_valves)
		}
	}
	fmt.Println("opened", IDstoValves(opened))

	return humanpoints + elephantpoints
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

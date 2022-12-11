package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"
	"sort"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 11

type Monkey struct {
	itemslist      []int
	itemscount     int
	operation      func(x int) int
	throwTo        func(x int) int
	inspectedCount int
}

func getOperation(s string) func(x int) int {
	operation := strings.Split(strings.Split(s, " = ")[1], " ")
	if operation[2] == "old" {
		if operation[1] == "*" {
			return func(x int) int { return x * x }
		} else {
			return func(x int) int { return x + x }
		}
	} else {
		c := inputs.ParseDecInt(operation[2])
		if operation[1] == "*" {
			return func(x int) int { return x * c }
		} else {
			return func(x int) int { return x + c }
		}
	}
}

func getThrower(ss []string) func(x int) int {
	divop, throwop1, throwop2 := strings.Split(ss[0], " "), strings.Split(ss[1], " "), strings.Split(ss[2], " ")
	divisor := inputs.ParseDecInt(divop[len(divop)-1])
	throw1 := inputs.ParseDecInt(throwop1[len(throwop1)-1])
	throw2 := inputs.ParseDecInt(throwop2[len(throwop2)-1])
	return func(x int) int {
		if x%divisor == 0 {
			return throw1
		}
		return throw2
	}
}

func part1(input string) interface{} {
	monkeyDefinitions := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(monkeyDefinitions))
	for i := range monkeyDefinitions {
		monkeyDef := strings.Split(monkeyDefinitions[i], "\n")

		startingItems := inputs.StrListToIntList(strings.Split(strings.Join(strings.Split(monkeyDef[1], " ")[4:], " "), ","))

		itemsList := make([]int, 1024)
		for i := range startingItems {
			itemsList[i] = startingItems[i]
		}

		monkeys[i] = Monkey{
			itemsList,
			len(startingItems),
			getOperation(monkeyDef[2]),
			getThrower(monkeyDef[3:]),
			0,
		}
	}

	fmt.Println(len(monkeys))

	for round := 0; round < 20; round++ {
		// fmt.Printf("Round %d\n", round)
		for i := range monkeys {
			// fmt.Printf("Working with Monkey %d\n", i)
			monkeyItems := monkeys[i].itemscount
			for j := 0; j < monkeyItems; j++ {
				// fmt.Printf("Monkey %d inspects %d\n", i, curr.Value)
				monkeys[i].inspectedCount++

				// operation
				monkeys[i].itemslist[j] = monkeys[i].operation(monkeys[i].itemslist[j])
				// fmt.Printf("\tWorry Level now %d\n", curr.Value)
				monkeys[i].itemslist[j] = monkeys[i].itemslist[j] / 3
				// fmt.Printf("\tMonkey bored Level now %d\n", curr.Value)
				throwto := monkeys[i].throwTo(monkeys[i].itemslist[j])
				// fmt.Printf("\tMonkey throws item %d to %d\n", curr.Value, throwto)
				monkeys[throwto].itemslist[monkeys[throwto].itemscount] = monkeys[i].itemslist[j]
				monkeys[i].itemscount--
				monkeys[throwto].itemscount++
			}
		}
	}

	return monkeyBusiness(monkeys)
}

func monkeyBusiness(monkeys []Monkey) int {
	monkeycounts := make([]int, len(monkeys))
	for i := range monkeys {
		monkeycounts[i] = monkeys[i].inspectedCount
	}
	sort.Ints(monkeycounts)
	return monkeycounts[len(monkeycounts)-1] * monkeycounts[len(monkeycounts)-2]
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

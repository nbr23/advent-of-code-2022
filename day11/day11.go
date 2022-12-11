package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

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

func getThrower(ss []string) (func(x int) int, int) {
	divop, throwop1, throwop2 := strings.Split(ss[0], " "), strings.Split(ss[1], " "), strings.Split(ss[2], " ")
	divisor := inputs.ParseDecInt(divop[len(divop)-1])
	throw1 := inputs.ParseDecInt(throwop1[len(throwop1)-1])
	throw2 := inputs.ParseDecInt(throwop2[len(throwop2)-1])
	return func(x int) int {
		if x%divisor == 0 {
			return throw1
		}
		return throw2
	}, divisor
}

func part1(input string) interface{} {
	monkeyDefinitions := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(monkeyDefinitions))
	for i := range monkeyDefinitions {
		monkeyDef := strings.Split(monkeyDefinitions[i], "\n")

		startingItems := inputs.StrListToIntList(strings.Split(strings.Join(strings.Split(monkeyDef[1], " ")[4:], " "), ","))

		itemsList := make([]int, 32)
		for i := range startingItems {
			itemsList[i] = startingItems[i]
		}

		thrower, _ := getThrower(monkeyDef[3:])

		monkeys[i] = Monkey{
			itemsList,
			len(startingItems),
			getOperation(monkeyDef[2]),
			thrower,
			0,
		}
	}

	for round := 0; round < 20; round++ {
		for i := range monkeys {
			monkeyItems := monkeys[i].itemscount
			for j := 0; j < monkeyItems; j++ {
				monkeys[i].inspectedCount++

				monkeys[i].itemslist[j] = monkeys[i].operation(monkeys[i].itemslist[j])
				monkeys[i].itemslist[j] = monkeys[i].itemslist[j] / 3
				throwto := monkeys[i].throwTo(monkeys[i].itemslist[j])
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
	divisors := 1
	monkeyDefinitions := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(monkeyDefinitions))
	for i := range monkeyDefinitions {
		monkeyDef := strings.Split(monkeyDefinitions[i], "\n")

		startingItems := inputs.StrListToIntList(strings.Split(strings.Join(strings.Split(monkeyDef[1], " ")[4:], " "), ","))

		itemsList := make([]int, 32)
		for i := range startingItems {
			itemsList[i] = startingItems[i]
		}

		thrower, divisor := getThrower(monkeyDef[3:])
		divisors *= divisor

		monkeys[i] = Monkey{
			itemsList,
			len(startingItems),
			getOperation(monkeyDef[2]),
			thrower,
			0,
		}
	}

	for round := 0; round < 10000; round++ {
		for i := range monkeys {
			monkeyItems := monkeys[i].itemscount
			for j := 0; j < monkeyItems; j++ {
				monkeys[i].inspectedCount++

				monkeys[i].itemslist[j] = monkeys[i].operation(monkeys[i].itemslist[j])
				if monkeys[i].itemslist[j] > divisors {
					monkeys[i].itemslist[j] = monkeys[i].itemslist[j] % divisors
				}
				throwto := monkeys[i].throwTo(monkeys[i].itemslist[j])
				monkeys[throwto].itemslist[monkeys[throwto].itemscount] = monkeys[i].itemslist[j]
				monkeys[i].itemscount--
				monkeys[throwto].itemscount++
			}
		}
	}
	return monkeyBusiness(monkeys)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 5

func parseStacks(input []string) []stack.Stack {
	stacks_count := len(strings.Split(input[len(input)-1], "  "))
	stacks := make([]stack.Stack, stacks_count)
	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < stacks_count; j++ {
			current := input[i][4*j : 4*j+3]
			if current != "   " {
				stacks[j].Push(current)
			}
		}
	}
	return stacks
}

func printStack(ss []stack.Stack) {
	for _, s := range ss {
		for s.Len() > 0 {
			fmt.Printf("%s ", s.Pop())
		}
		fmt.Println()
	}
}

func makeMove(move string, ss []stack.Stack) {
	spl := strings.Split(move, " ")
	count := inputs.ParseDecInt(spl[1])
	src := inputs.ParseDecInt(spl[3]) - 1
	dest := inputs.ParseDecInt(spl[5]) - 1

	for i := 0; i < count; i++ {
		c := ss[src].Pop()
		ss[dest].Push(c)
	}
}

func part1(input string) interface{} {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(strings.Split(parts[0], "\n"))

	for _, move := range inputs.InputToStrList(parts[1]) {
		makeMove(move, stacks)
	}

	res := ""
	for _, s := range stacks {
		p := fmt.Sprintf("%v", s.Peek())
		res = fmt.Sprintf("%s%c", res, p[1])
	}
	return res
}

func makeMove2(move string, ss []stack.Stack) {
	spl := strings.Split(move, " ")
	count := inputs.ParseDecInt(spl[1])
	src := inputs.ParseDecInt(spl[3]) - 1
	dest := inputs.ParseDecInt(spl[5]) - 1

	var tmp stack.Stack

	for i := 0; i < count; i++ {
		c := ss[src].Pop()
		tmp.Push(c)
	}
	for i := 0; i < count; i++ {
		c := tmp.Pop()
		ss[dest].Push(c)
	}
}

func part2(input string) interface{} {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(strings.Split(parts[0], "\n"))

	for _, move := range inputs.InputToStrList(parts[1]) {
		makeMove2(move, stacks)
	}

	res := ""
	for _, s := range stacks {
		p := fmt.Sprintf("%v", s.Peek())
		res = fmt.Sprintf("%s%c", res, p[1])
	}
	return res
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

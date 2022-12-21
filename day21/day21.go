package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"
	"strconv"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 21

func compute(data map[string]string, key string) int {
	num, err := strconv.ParseInt(strings.Trim(data[key], " "), 10, 64)

	if err == nil {
		return int(num)
	}
	parts := strings.Split(data[key], " ")
	op := parts[1]
	if op == "+" {
		return compute(data, parts[0]) + compute(data, parts[2])
	}

	if op == "-" {
		return compute(data, parts[0]) - compute(data, parts[2])
	}

	if op == "/" {
		return compute(data, parts[0]) / compute(data, parts[2])
	}

	if op == "*" {
		return compute(data, parts[0]) * compute(data, parts[2])
	}
	panic("never happens")
}

func compute3(data map[string]string, key string) func(value int64) int64 {
	if key == "humn" {
		return func(value int64) int64 { return value }
	}

	num, err := strconv.ParseInt(strings.Trim(data[key], " "), 10, 64)

	if err == nil {
		return func(value int64) int64 { return int64(num) }
	}

	parts := strings.Split(data[key], " ")
	op := parts[1]
	if op == "+" {
		return func(value int64) int64 { return compute3(data, parts[0])(value) + compute3(data, parts[2])(value) }
	}

	if op == "-" {
		return func(value int64) int64 { return compute3(data, parts[0])(value) - compute3(data, parts[2])(value) }
	}

	if op == "/" {
		return func(value int64) int64 { return compute3(data, parts[0])(value) / compute3(data, parts[2])(value) }
	}

	if op == "*" {
		return func(value int64) int64 { return compute3(data, parts[0])(value) * compute3(data, parts[2])(value) }
	}
	panic("never happens")
}

func part1(input string) interface{} {
	inputList := inputs.InputToStrList(input)
	data := make(map[string]string)
	for _, line := range inputList {
		parts := strings.Split(line, ": ")
		hash := parts[0]
		data[hash] = parts[1]
	}

	return compute(data, "root")
}

func part2(input string) interface{} {
	inputList := inputs.InputToStrList(input)
	data := make(map[string]string)
	for _, line := range inputList {

		parts := strings.Split(line, ": ")
		hash := parts[0]
		data[hash] = parts[1]

	}
	parts := strings.Split(data["root"], " ")
	c1 := compute3(data, parts[0])
	c2 := compute3(data, parts[2])
	for i := int64(3876027194999); ; i++ { // manual dichotomy done here...
		if i%1000 == 0 {
			fmt.Println(i, ": ", c1(i), "==", c2(i))
		}
		if c1(i) == c2(i) {
			return i
		}
	}
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

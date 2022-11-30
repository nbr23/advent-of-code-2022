package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
)

func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type Resolver func(string) interface{}

func Solve(part1 Resolver, part2 Resolver, day int) {
	var input string
	token_arg := flag.String("token", "", "Authentication token to retrieve the puzzle input. If not set, will look for it in the file ./.token")
	force_fetch_input := flag.Bool("fetch_input", false, "Force fetching of the input. Will overwrite existing input file")
	flag.Parse()

	binput, err := os.ReadFile(fmt.Sprintf("./inputs/day%02d.txt", day))

	if err != nil || *force_fetch_input {
		var token string
		fmt.Println("Fetching input")
		if token_arg == nil || len(*token_arg) == 0 {
			token = inputs.GetToken()
		} else {
			token = *token_arg
		}
		input = inputs.GetInput(day, token)
		err := os.WriteFile(fmt.Sprintf("./inputs/day%02d.txt", day), []byte(input), 0700)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		input = string(binput)
	}

	fmt.Printf("*** DAY %d ***\n", day)

	res1 := part1(input)
	fmt.Printf("Result 1: %v\n", res1)

	res2 := part2(input)

	fmt.Printf("Result 2: %v\n", res2)
}

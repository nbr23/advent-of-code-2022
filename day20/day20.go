package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"fmt"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/linkedlist"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 20

func findZero(l *linkedlist.LinkedList[int]) *linkedlist.LinkedList[int] {
	if l.Value == 0 {
		return l
	}
	return findZero(l.Next)
}

func part1(input string) interface{} {
	input_i := inputs.InputToIntList(input)
	inputList, elements := linkedlist.ToLinkedList(input_i)
	linkedlist.MakeCircular(inputList)

	for _, e := range elements {
		if e.Value == 0 {
			continue
		}
		// We remove e from its current position
		old_prev := e.Previous
		old_next := e.Next
		old_prev.Next = old_next
		old_next.Previous = old_prev
		if e.Value > 0 {
			current := e
			// move right for e.Value steps
			for i := 0; i < e.Value; i++ {
				current = current.Next
			}

			// save the current next and previous
			oldnext := current.Next

			// insert e
			current.Next = e
			e.Previous = current
			oldnext.Previous = e
			e.Next = oldnext
		} else {
			current := e
			// move right for e.Value steps
			for i := 0; i > e.Value; i-- {
				current = current.Previous
			}

			// save the current next and previous
			oldprev := current.Previous

			// insert e
			current.Previous = e
			e.Next = current
			oldprev.Next = e
			e.Previous = oldprev

		}
	}

	list := findZero(inputList).ToList()
	fmt.Println(list)

	return list[1000%len(list)] + list[2000%len(list)] + list[3000%len(list)]
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

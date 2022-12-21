package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/linkedlist"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 20

func mod(i, n int) int {
	return ((i % n) + n) % n
}

// find out who is before and after a node
func getBeforeAfter[T any](l *linkedlist.LinkedList[T], n *linkedlist.Node[T]) (*linkedlist.Node[T], *linkedlist.Node[T]) {
	prev := l.Head
	curr := l.Head.Next
	for {
		if curr == n {
			return prev, curr.Next
		}
		prev = curr
		curr = curr.Next
	}
}

// resturns the Node right before and right after the new spot
func findNewSpot[T any](n *linkedlist.Node[T], shifts int) (*linkedlist.Node[T], *linkedlist.Node[T]) {
	curr := n
	for i := 0; i < shifts; i++ {
		curr = curr.Next
	}
	return curr, curr.Next
}

func part1(input string) interface{} {
	return compute(input, 1, 1)
}
func part2(input string) interface{} {
	return compute(input, 10, 811589153)
}

func compute(input string, shuffles int, key int) interface{} {
	SHUFFLES := shuffles
	input_i := inputs.InputToIntList(input)
	for i := range input_i {
		input_i[i] = input_i[i] * key
	}
	linkedList, nodes := linkedlist.ToLinkedList(input_i)

	for ; SHUFFLES > 0; SHUFFLES-- {
		for _, e := range nodes {
			// No action on 0
			if mod(e.Value, len(input_i)-1) == 0 {
				continue
			}

			// How many shifts do we need to do
			shifts := mod(e.Value, len(input_i)-1)

			// Find out where we are now
			before_me, after_me := getBeforeAfter(linkedList, e)

			// find the new place
			before, after := findNewSpot(e, shifts)

			// insert at new place
			before.Next = e
			e.Next = after

			// remove from the old place
			before_me.Next = after_me
		}
	}
	linkedList.SetHeadToZero()
	intlist := linkedList.ToList()
	return intlist[1000%len(intlist)] + intlist[2000%len(intlist)] + intlist[3000%len(intlist)]
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

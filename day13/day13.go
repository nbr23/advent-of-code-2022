package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"
	"fmt"
	"strconv"
	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 13

type Item struct {
	IsList  bool
	IntVal  int64
	ListVal []*Item
}

func (i Item) String() string {
	if !i.IsList {
		return fmt.Sprintf("%d", i.IntVal)
	}
	s := "["
	for _, v := range i.ListVal {
		s = fmt.Sprintf("%s%s,", s, v.String())
	}
	s = fmt.Sprintf("%s]", s)
	return s
}

func splitPacket(input string) *Item {
	if len(input) == 0 {
		return nil
	}
	if i, err := strconv.ParseInt(strings.Trim(input, " "), 10, 64); err == nil {
		return &Item{false, i, nil}
	}

	input = input[1 : len(input)-1]
	item := Item{true, 0, []*Item{}}

	current := ""
	level := 0
	for _, c := range input {
		if c == '[' {
			level++
			current = fmt.Sprintf("%s%c", current, c)
			continue
		} else if c == ']' {
			level--
			current = fmt.Sprintf("%s%c", current, c)
			continue
		}
		if c == ',' && level == 0 {
			item.ListVal = append(item.ListVal, splitPacket(current))
			current = ""
		} else {
			current = fmt.Sprintf("%s%c", current, c)
		}
	}
	item.ListVal = append(item.ListVal, splitPacket(current))
	return &item
}

func compare(p1, p2 *Item) int {
	if p1 == nil {
		if p2 == nil {
			return 0
		}
		return 1
	} else if p2 == nil {
		return -1
	}
	if !p1.IsList && !p2.IsList {
		if p1.IntVal == p2.IntVal {
			return 0
		}
		if p1.IntVal > p2.IntVal {
			return -1
		}
		return 1
	}

	if !p1.IsList {
		p1 = &Item{true, 0, []*Item{p1}}
	}
	if !p2.IsList {
		p2 = &Item{true, 0, []*Item{p2}}
	}
	size := len(p1.ListVal)
	if len(p2.ListVal) < len(p1.ListVal) {
		size = len(p2.ListVal)
	}

	for i := 0; i < size; i++ {
		c := compare(p1.ListVal[i], p2.ListVal[i])
		if c == -1 {
			return -1
		} else if c == 1 {
			return 1
		}
	}
	if len(p1.ListVal) < len(p2.ListVal) {
		return 1
	} else if len(p1.ListVal) > len(p2.ListVal) {
		return -1
	}
	return 0
}

func part1(input string) interface{} {
	pairs := strings.Split(input, "\n\n")
	res := 0

	for i, pair := range pairs {
		packets := strings.Split(pair, "\n")
		p1 := splitPacket(packets[0])
		p2 := splitPacket(packets[1])
		if compare(p1, p2) == 1 {
			res += i + 1
		}
	}
	return res
}

type LinkedList struct {
	Next  *LinkedList
	Value *Item
}

func Append(l *LinkedList, v *Item, acc int) (*LinkedList, int) {
	if l == nil {
		return &LinkedList{nil, v}, acc
	}

	if l.Next == nil {
		if compare(l.Value, v) < 1 {
			newl := &LinkedList{l, v}
			return newl, acc + 1
		}
		l.Next = &LinkedList{nil, v}
		return l, acc + 1
	}

	curr := l
	for {
		if curr.Next == nil {
			curr.Next = &LinkedList{nil, v}
			break
		}
		if compare(curr.Next.Value, v) < 1 {
			newl := &LinkedList{curr.Next, v}
			curr.Next = newl
			return l, acc + 1
		}
		curr = curr.Next
		acc++
	}
	return l, acc + 1
}

func (l *LinkedList) String() string {
	if l == nil {
		return "nil"
	}
	v := ""
	curr := l
	for {
		if curr.Value == nil {
			v = fmt.Sprintf("%s[] ", v)
		} else {
			v = fmt.Sprintf("%s%v\n", v, curr.Value)
		}
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	return fmt.Sprintf("%s\n", v)
}

func part2(input string) interface{} {
	pairs := strings.Split(input, "\n\n")
	var ll *LinkedList

	for _, pair := range pairs {
		packets := strings.Split(pair, "\n")
		p1 := splitPacket(packets[0])
		p2 := splitPacket(packets[1])
		ll, _ = Append(ll, p1, 0)
		ll, _ = Append(ll, p2, 0)
	}
	ll, a := Append(ll, splitPacket("[[2]]"), 0)
	ll, b := Append(ll, splitPacket("[[6]]"), 0)
	return (a + 1) * (b + 1)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}

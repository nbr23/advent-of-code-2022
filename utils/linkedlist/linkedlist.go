package linkedlist

import (
	"fmt"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (n Node[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) SetHeadToZero() {
	for {
		if inputs.ParseDecInt(fmt.Sprint(l.Head.Value)) == 0 { // lol lol
			return
		}
		l.Head = l.Head.Next
	}
}

func (l *LinkedList[T]) Append(v T) {
	if l.Tail == nil && l.Head == nil {
		l.Tail = &Node[T]{v, l.Head}
		l.Head = l.Tail
	} else {
		l.Tail.Next = &Node[T]{v, l.Head}
		l.Tail = l.Tail.Next
	}
}

func (l *LinkedList[T]) ToList() []T {
	if l.Head == nil && l.Tail == nil {
		return []T{}
	}
	curr := l.Head
	res := make([]T, 0)
	for {
		res = append(res, curr.Value)
		curr = curr.Next
		if curr == l.Head {
			break
		}
	}
	return res
}

func ToLinkedList[T any](l []T) (*LinkedList[T], []*Node[T]) {
	ll := &LinkedList[T]{}
	nodes := make([]*Node[T], len(l))

	if len(l) == 0 {
		return nil, nodes
	}
	for i, e := range l {
		ll.Append(e)
		nodes[i] = ll.Tail
	}
	return ll, nodes
}

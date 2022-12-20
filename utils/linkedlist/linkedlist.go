package linkedlist

import "fmt"

type LinkedList[T any] struct {
	Previous *LinkedList[T]
	Next     *LinkedList[T]
	Value    T
}

// (newlist, element)
func Append[T any](l *LinkedList[T], v T) (*LinkedList[T], *LinkedList[T]) {
	var newelt *LinkedList[T]
	if l == nil {
		newelt = &LinkedList[T]{nil, nil, v}
		return newelt, newelt
	}
	curr := l
	for {
		if curr.Next == nil {
			curr.Next = &LinkedList[T]{curr, nil, v}
			newelt = curr.Next
			break
		}
		curr = curr.Next
	}
	return l, newelt
}

func PopLast[T any](l *LinkedList[T]) (*LinkedList[T], T) {
	var v T
	curr := l
	prev := l
	if curr.Next == nil {
		return nil, curr.Value
	}
	for {
		if curr.Next == nil {
			prev.Next = nil
			v = curr.Value
			break
		}
		prev = curr
		curr = curr.Next
	}
	return l, v
}

func ToLinkedList[T any](l []T) (*LinkedList[T], []*LinkedList[T]) {
	var ll *LinkedList[T]
	var newelt *LinkedList[T]
	nodes := make([]*LinkedList[T], len(l))
	if len(l) == 0 {
		return nil, nodes
	}
	for i := range l {
		ll, newelt = Append(ll, l[i])
		nodes[i] = newelt
	}
	return ll, nodes
}

func MakeCircular[T any](l *LinkedList[T]) {
	if l == nil {
		return
	}
	first := l
	for {
		if l.Next == nil {
			l.Next = first
			first.Previous = l
			return
		}
		l = l.Next
	}
}

func (l *LinkedList[T]) ToList() []T {
	resList := make([]T, 0)
	if l == nil {
		return resList
	}
	curr := l
	looped := false
	for {
		if curr == l {
			if looped {
				break
			}
			looped = true
		}
		resList = append(resList, curr.Value)
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	return resList
}

func (l *LinkedList[T]) String() string {
	if l == nil {
		return "nil"
	}
	v := "[ "
	curr := l
	looped := false
	for {
		if curr == l {
			if looped {
				break
			}
			looped = true
		}
		v = fmt.Sprintf("%s%v ", v, curr.Value)
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	return fmt.Sprintf("%s]", v)
}

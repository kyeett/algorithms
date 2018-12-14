package stacks

import "fmt"

// LinkedStack implements a stack using a linked list
type LinkedStack struct {
	first  *node
	format int
}

// Push item to stack
func (ls *LinkedStack) Push(item string) {
	ls.first = &node{
		val:  item,
		next: ls.first,
	}
}

// Pop item from stack
func (ls *LinkedStack) Pop() string {
	n := ls.first
	ls.first = ls.first.next
	return n.val
}

// Iter returns a channel that can be iterated over
func (ls *LinkedStack) Iter() <-chan string {
	ch := make(chan string)
	go func() {
		for !ls.isEmpty() {
			ch <- ls.Pop()
		}
		close(ch)
	}()
	return ch
}

func (ls LinkedStack) isEmpty() bool {
	return ls.first == nil
}

func (ls LinkedStack) String() string {
	// format := "%d"
	// if ls.format > 0 {
	// 	format = fmt.Sprintf("[%% %d%%d]", ls.format)
	// }

	n := ls.first
	s := ""
	for n != nil {
		s += fmt.Sprintf(" % 4s", n.val)
		n = n.next
	}
	return s
}

type node struct {
	val  string
	next *node
}

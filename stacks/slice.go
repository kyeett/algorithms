package stacks

import "fmt"

// SliceStack implements a stack using the Go slice
type SliceStack struct {
	stack []string
}

// Push item to stack
func (as *SliceStack) Push(item string) {
	as.stack = append(as.stack, item)
}

// Pop item from stack
func (as *SliceStack) Pop() string {
	item := as.stack[len(as.stack)-1]
	as.stack = as.stack[:len(as.stack)-1]
	return item
}

func (as SliceStack) isEmpty() bool {
	return len(as.stack) == 0
}

// Iter returns a channel that can be iterated over
func (as *SliceStack) Iter() <-chan string {
	ch := make(chan string)
	go func() {
		for !as.isEmpty() {
			ch <- as.Pop()
		}
		close(ch)
	}()
	return ch
}

func (as SliceStack) String() string {
	s := ""
	// s += fmt.Sprintf(" (%d,%d) ", as.n, as.maxN)
	for _, i := range as.stack {
		s += fmt.Sprintf("% -4s", i)
	}
	for i := 0; i < cap(as.stack)-len(as.stack); i++ {
		s += fmt.Sprintf("% -4s", ".")
	}

	return s
}

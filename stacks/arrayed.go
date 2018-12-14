package stacks

import "fmt"

const maxN = 7

// ArrayStack implements a stack using a fixed length array
type ArrayStack struct {
	n     int
	stack [maxN]string
}

// Push item to stack
func (as *ArrayStack) Push(item string) {
	as.stack[as.n] = item
	as.n++
}

// Pop item from stack
func (as *ArrayStack) Pop() string {
	as.n--
	return as.stack[as.n]
}

func (as ArrayStack) isEmpty() bool {
	return as.n == 0
}

// Iter returns a channel that can be iterated over
func (as *ArrayStack) Iter() <-chan string {
	ch := make(chan string)
	go func() {
		for !as.isEmpty() {
			ch <- as.Pop()
		}
		close(ch)
	}()
	return ch
}

func (as ArrayStack) String() string {
	s := ""
	// s += fmt.Sprintf("(%d,%d) ", as.n, maxN)
	for i := 0; i < as.n; i++ {
		s += fmt.Sprintf("% -5s", as.stack[as.n-i-1])
	}
	for i := 0; i < maxN-as.n; i++ {
		s += fmt.Sprintf("% -5s", ".")
	}

	return s
}

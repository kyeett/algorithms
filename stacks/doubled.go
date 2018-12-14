package stacks

import "fmt"

// DoublingStack implements a stack with variable length (a worse version of the Go slice, for training purposes :-)
type DoublingStack struct {
	n     int
	maxN  int
	stack []string
}

func (as *DoublingStack) resize(ln int) {
	new := make([]string, ln)
	copy(new, as.stack)
	as.stack = new
	as.maxN = ln
}

// Push item to stack
func (as *DoublingStack) Push(item string) {
	if as.n+1 > as.maxN {
		as.resize(2 * (as.n + 1))
	}
	as.stack[as.n] = item
	as.n++
}

// Pop item from stack
func (as *DoublingStack) Pop() string {
	as.n--
	if as.n < as.maxN/4 {
		as.resize(as.maxN / 2)
	}
	return as.stack[as.n]
}

func (as DoublingStack) isEmpty() bool {
	return as.n == 0
}

// Iter returns a channel that can be iterated over
func (as *DoublingStack) Iter() <-chan string {
	ch := make(chan string)
	go func() {
		for !as.isEmpty() {
			ch <- as.Pop()
		}
		close(ch)
	}()
	return ch
}

func (as DoublingStack) String() string {
	s := ""
	// s += fmt.Sprintf(" (%d,%d) ", as.n, as.maxN)
	for i := 0; i < as.n; i++ {
		s += fmt.Sprintf("% -5s", as.stack[as.n-i-1])
	}
	for i := 0; i < as.maxN-as.n; i++ {
		s += fmt.Sprintf("% -5s", ".")
	}

	return s
}

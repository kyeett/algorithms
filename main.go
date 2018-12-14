package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/kyeett/algorithms/stacks"
)

func main() {
	fmt.Println("Examples of stacks")
	fmt.Println("")

	input, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	ls := stacks.LinkedStack{}
	as := stacks.ArrayStack{}
	for _, op := range strings.Fields(string(input)) {
		switch op {
		case "-":
			ls.Pop()
			as.Pop()
		default:
			ls.Push(op)
			as.Push(op)
		}
		fmt.Printf("% 40s% 40s\n", ls, as)
	}

	fmt.Println("Flush stacks")
	lsCh := ls.Iter()
	asCh := as.Iter()
	for i := 0; i < 5; i++ {
		fmt.Printf("% 40s% 40s\n", <-lsCh, <-asCh)
	}
}

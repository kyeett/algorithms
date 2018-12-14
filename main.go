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
	ds := stacks.DoublingStack{}
	for _, op := range strings.Fields(string(input)) {
		switch op {
		case "-":
			ls.Pop()
			as.Pop()
			ds.Pop()
		default:
			ls.Push(op)
			as.Push(op)
			ds.Push(op)
		}
		fmt.Printf("% -30s| % -35s| % -35s\n", ls, as, ds)
	}

	fmt.Println("\nFlush stacks")
	lsCh := ls.Iter()
	asCh := as.Iter()
	dsCh := ds.Iter()
	for i := 0; i < 5; i++ {
		fmt.Printf("% -30s| % -35s| % -35s\n", <-lsCh, <-asCh, <-dsCh)
	}
}

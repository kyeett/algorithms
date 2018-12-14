package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/kyeett/algorithms/stacks"
)

func main() {
	fmt.Println("Examples of stacks")
	fmt.Println("")

	input, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	format := "% -30s| % -30s| % -30s| % -30s\n"

	ls := stacks.LinkedStack{}
	as := stacks.ArrayStack{}
	ds := stacks.DoublingStack{}
	ss := stacks.SliceStack{}
	fmt.Printf(format, "LinkedStack", "ArrayStack", "DoublingStack", "SliceStack")
	for _, op := range strings.Fields(string(input)) {
		switch op {
		case "-":
			ls.Pop()
			as.Pop()
			ds.Pop()
			ss.Pop()
		default:
			ls.Push(op)
			as.Push(op)
			ds.Push(op)
			ss.Push(op)
		}
		fmt.Printf(format, ls, as, ds, ss)
	}

	fmt.Println("\nFlush stacks")
	lsCh := ls.Iter()
	asCh := as.Iter()
	dsCh := ds.Iter()
	ssCh := ss.Iter()
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond) //Workaround for aestatic purposes, to keep all chans in check :-)
		fmt.Printf(format, ls, as, ds, ss)
		time.Sleep(10 * time.Millisecond)
		<-lsCh
		<-asCh
		<-dsCh
		<-ssCh
	}
}

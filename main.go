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
	for _, op := range strings.Fields(string(input)) {
		switch op {
		case "-":
			ls.Pop()
		default:
			ls.Push(op)
		}
		fmt.Println(ls)
	}

	for s := range ls.Iter() {
		fmt.Println(s)
	}
}

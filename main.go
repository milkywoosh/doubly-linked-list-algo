package main

import "fmt"

var l *LinkedList = &LinkedList{parent: 10}

func main() {

	l.Add(100)
	l.Add(30)
	l.Add(50)
	l.Add(19)
	l.Add(48)

	l.Print()

	l = l.Prepend(300)
	fmt.Print("limit ============\n")
	fmt.Print("limit ============\n")
	l.Print()
}

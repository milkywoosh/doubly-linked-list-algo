package main

import "fmt"

// hint of size = 100, avoid rehashing and avoid more memory allocation and CPU cost
var hm HashMap = make(HashMap, 3)

func main() {

	// fmt.Printf("test %s", "Last Recently Used")

	// hm[1] = 1000
	// hm[2] = 2000
	// hm[3] = 4000
	// hm[4] = 5000
	// hm[5] = 6000

	// fmt.Printf("hm %v", hm)
	// fmt.Printf("check cap: %v", len(hm))

	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	initNode := NewDoublyLinkedList()

	fmt.Printf("head: %v\n", initNode.head)
	fmt.Printf("head prev: %v\n", initNode.head.prev)
	fmt.Printf("head next: %v\n", initNode.head.next)

	fmt.Printf("tail: %v\n", initNode.tail)
	fmt.Printf("tail prev: %v\n", initNode.tail.prev)
	fmt.Printf("tail next: %v\n", initNode.tail.next)
}

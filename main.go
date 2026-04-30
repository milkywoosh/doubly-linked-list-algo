package main

import (
	"fmt"
	"log"

	dllhash "github.com/structgo/dll_hash"
)

// hint of size = 100, avoid rehashing and avoid more memory allocation and CPU cost

// const (
// 	SmallVehicle  VehicleSize = iota
// 	MediumVehicle VehicleSize = iota
// 	BigVehicle    VehicleSize = iota
// )

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("err recover: %v", err)
			// return
		}
	}()

	initNode := NewDoublyLinkedList()

	initNode.AddToHead(3)
	initNode.AddToHead(100)
	initNode.AddToTail(10)
	initNode.AddToTail(4)
	initNode.AddToTail(3)
	initNode.AddToTail(8)
	initNode.AddToHead(89)

	initNode.TraverseToTail()
	initNode.TraverseToHead()

	initHashDll := dllhash.NewDoublyLinkedList()

	initHashDll.AddToTail(76)
	initHashDll.AddToHead(10)
	initHashDll.AddToHead(32)
	initHashDll.AddToTail(94)

	_, err := initHashDll.InfoCurrNode(76)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// fmt.Printf("curr Node: %v\n", *curr)
	// fmt.Printf("curr val: %v\n", curr.Val)
	// fmt.Printf("curr prev: %v\n", curr.Prev)
	// fmt.Printf("curr prev.Val: %v\n", curr.Prev.Val)
	// fmt.Printf("curr next: %v\n", curr.Next)
	// fmt.Printf("curr next.val: %v\n\n", curr.Next.Val)

	err = initHashDll.AddPrev(94, 66)
	if err != nil {
		fmt.Printf("err AddPrev: %v", err)
		return
	}
	err = initHashDll.AddPrev(10, 44)
	if err != nil {
		fmt.Printf("err AddPrev: %v", err)
		return
	}

	// initHashDll.TraverseToTail() // 32, 10, 76, 66, 94

}

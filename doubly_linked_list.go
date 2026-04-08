package main

import "fmt"

type HashMap map[int]int

// note that DoublyLinkedListNode doesnt know anything
type DoublyLinkedListNode struct {
	val        int
	next, prev *DoublyLinkedListNode
}

func newNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		val:  val,
		next: nil,
		prev: nil,
	}
}

// wrapper, this struct drive the method
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	headNode := &DoublyLinkedListNode{-1, nil, nil}
	tailNode := &DoublyLinkedListNode{10000, nil, nil}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &DoublyLinkedList{
		head: headNode,
		tail: tailNode,
	}
}

func (d *DoublyLinkedList) AddTail(val int) {

	newNode := newNode(val)
	if d.tail == nil {
		d.tail = newNode
		d.head = newNode
		return
	}

	// move tail to before new node
	newNode.prev = d.tail
	newNode.next = nil

	d.tail.next = newNode
	d.tail = newNode
}

func (d *DoublyLinkedList) TraverseToHead() {
	fmt.Printf("only d: %v\n", d)

	if d != nil {
		for d.tail != nil {
			fmt.Printf("tail ...: %v\n", d.tail.val)

			d.tail = d.tail.prev
		}
	}
	if d.tail == nil {
		fmt.Printf("data tail is: %v\n", d)
		return
	}

	for d.tail.next != nil {
		fmt.Printf("next to tail next: %v\n", d.tail)
		d.tail = d.tail.next
	}
}

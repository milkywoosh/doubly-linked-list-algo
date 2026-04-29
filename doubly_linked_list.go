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
	// headNode := &DoublyLinkedListNode{val: -1, prev: nil, next: nil}
	// tailNode := &DoublyLinkedListNode{val: -1, prev: nil, next: nil}

	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) AddToTail(val int) {
	newNode := newNode(val)
	if d.tail == nil {
		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		// so initially, or at the first addition, newNode is actually d.tail and d.head at the same time
		d.tail = newNode
		d.head = newNode
	} else {

		newNode.prev = d.tail
		// d.tail saat ini masih memory address node awal
		d.tail.next = newNode
		d.tail = newNode
	}
}

func (d *DoublyLinkedList) AddToHead(val int) {
	newNode := newNode(val)

	if d.head == nil {
		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		// so initially, or at the first addition, newNode is actually d.tail and d.head at the same time
		d.tail = newNode
		d.head = newNode

	} else {

		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		newNode.next = d.head
		fmt.Println(newNode.next == d.head)
		d.head.prev = newNode
		d.head = newNode
		fmt.Println(newNode.next == d.head)

	}

}

func (d *DoublyLinkedList) TraverseToHead() {
	fmt.Printf("only d: %v\n", d)

	if d != nil {
		for d.tail != nil {
			fmt.Printf("tail ...: %v\n", d.tail.val)
			d.tail = d.tail.prev
		}
	}

}

func (d *DoublyLinkedList) TraverseToTail() {
	// note: dont forget to use TEMP variable to avoid changing to the REAL memory address
	tempNode := d.head
	if d != nil {
		for tempNode != nil {
			fmt.Printf("from head: %v\n", tempNode.val)

			tempNode = tempNode.next

		}
	}
}

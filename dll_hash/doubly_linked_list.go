package dllhash

import (
	"fmt"
	"sync"
)

type HashMap map[int]*DoublyLinkedListNode

// note that DoublyLinkedListNode doesnt know anything
type DoublyLinkedListNode struct {
	Val        int
	Next, Prev *DoublyLinkedListNode
}

func newNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

// wrapper, this struct drive the method
type DoublyLinkedList struct {
	mu    sync.Mutex
	hashd HashMap
	head  *DoublyLinkedListNode
	tail  *DoublyLinkedListNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	// headNode := &DoublyLinkedListNode{val: -1, prev: nil, next: nil}
	// tailNode := &DoublyLinkedListNode{val: -1, prev: nil, next: nil}

	return &DoublyLinkedList{
		hashd: make(HashMap),
		head:  nil,
		tail:  nil,
	}
}

func (d *DoublyLinkedList) SetHash(key int, node *DoublyLinkedListNode) {
	d.mu.Lock()
	d.hashd[key] = node
	d.mu.Unlock()
}

func (d *DoublyLinkedList) InfoCurrNode(key int) (*DoublyLinkedListNode, error) {
	// big O => O(1)
	if val, ok := d.hashd[key]; !ok {
		return val, fmt.Errorf("nilai berikut tidak ada : %d", key)
	}

	memAddr := d.hashd[key]
	return memAddr, nil
}

func (d *DoublyLinkedList) AddToTail(val int) {
	newNode := newNode(val)
	if d.tail == nil {
		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		// so initially, or at the first addition, newNode is actually d.tail and d.head at the same time
		d.tail = newNode
		d.head = newNode
		d.SetHash(val, newNode)
	} else {

		newNode.Prev = d.tail
		// d.tail saat ini masih memory address node awal
		d.tail.Next = newNode
		d.tail = newNode
		d.SetHash(val, newNode)
	}
}

func (d *DoublyLinkedList) AddToHead(val int) {
	newNode := newNode(val)

	if d.head == nil {
		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		// so initially, or at the first addition, newNode is actually d.tail and d.head at the same time
		d.tail = newNode
		d.head = newNode
		d.SetHash(val, newNode)

	} else {

		// note: tail and head, pointing to the same memory address newNode *DoublyLinkedListNode
		newNode.Next = d.head
		// fmt.Println(newNode.next == d.head)
		d.head.Prev = newNode
		d.head = newNode
		// fmt.Println(newNode.next == d.head)
		d.SetHash(val, newNode)

	}

}

func (d *DoublyLinkedList) TraverseToHead() {
	fmt.Printf("only d: %v\n", d)

	if d != nil {
		for d.tail != nil {
			fmt.Printf("tail ...: %v\n", d.tail.Val)
			d.tail = d.tail.Prev
		}
	}

}

func (d *DoublyLinkedList) TraverseToTail() {
	// note: dont forget to use TEMP variable to avoid changing to the REAL memory address
	tempNode := d.head
	if d != nil {
		for tempNode != nil {
			fmt.Printf("from head: %v\n", tempNode.Val)

			tempNode = tempNode.Next

		}
	}
}

// problem if add to previous of Head
func (d *DoublyLinkedList) AddBefore(currVal, val int) error {
	// var memAddr *DoublyLinkedListNode
	memAddr, err := d.InfoCurrNode(currVal)
	if err != nil {
		return fmt.Errorf("val %d tidak ditemukan", currVal)
	}

	node := newNode(val) // memory address
	d.SetHash(val, node)

	fmt.Println("curr val prev: ", memAddr.Prev)

	node.Next = memAddr
	fmt.Println("node.Next: ", node.Next)
	fmt.Println("node.Val: ", node.Val)
	node.Prev = memAddr.Prev
	fmt.Println("node.Prev: ", node.Prev)
	if memAddr.Prev != nil {
		memAddr.Prev.Next = node
	} else {

	}

	return nil
}

func (d *DoublyLinkedList) Find(val int) bool {
	_, ok := d.hashd[val]
	if !ok {
		return false
	} else {
		return ok
	}
}

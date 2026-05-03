package dllhash

import (
	"fmt"
	"sync"
)

type HashMap map[int]*DoublyLinkedListNode

//
// << HEAD (prev) =================== (next) TAIL >>

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

		d.head.Prev = newNode
		d.head = newNode

		d.SetHash(val, newNode)

	}

}

func (d *DoublyLinkedList) TraverseTailToHead() {
	tempNode := d.tail
	if d != nil {
		fmt.Printf("from tail to head \n")
		for tempNode != nil {
			fmt.Printf("%v - ", tempNode.Val)
			tempNode = tempNode.Prev
		}
	}

}

func (d *DoublyLinkedList) TraverseHeadToTail() {
	// note: dont forget to use TEMP variable to avoid changing to the REAL memory address
	tempNode := d.head
	if d != nil {
		fmt.Printf("from head to tail \n")
		for tempNode != nil {
			fmt.Printf("%v - ", tempNode.Val)

			tempNode = tempNode.Next

		}
	}
}

// problem if add to previous of Head
func (d *DoublyLinkedList) AddPrev(currVal, val int) error {
	// var memAddr *DoublyLinkedListNode
	memAddr, err := d.InfoCurrNode(currVal)
	if err != nil {
		return fmt.Errorf("val %d tidak ditemukan", currVal)
	}

	if memAddr.Prev != nil {
		node := newNode(val) // memory address

		memAddr.Prev.Next = node
		node.Next = memAddr
		node.Prev = memAddr.Prev
		memAddr.Prev = node

		d.SetHash(val, node)
		return nil
	} else {
		// if nil use existed API
		d.AddToHead(val)
		return nil

	}

}

func (d *DoublyLinkedList) Find(val int) bool {
	_, ok := d.hashd[val]
	return ok
}

func (d *DoublyLinkedList) GetHead() *DoublyLinkedListNode {
	if d != nil {
		if d.head != nil {
			return d.head
		} else {
			return nil
		}
	}
	return nil
}

func (d *DoublyLinkedList) GetTail() *DoublyLinkedListNode {
	if d != nil {
		if d.tail != nil {
			return d.tail
		} else {
			return nil
		}
	}
	return nil
}

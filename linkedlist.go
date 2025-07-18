package main

import "fmt"

type LinkedList struct {
	parent int
	child  *LinkedList
}

func (ll *LinkedList) Add(val int) {
	if ll.child == nil {
		ll.child = &LinkedList{parent: val}
		return
	} else {
		ll.child.Add(val)
		return
	}
}

func (ll *LinkedList) Print() {
	current := ll
	for current != nil {
		fmt.Printf("%d\n", current.parent)
		current = current.child
	}
}

func (ll *LinkedList) Prepend(val int) *LinkedList {
	// insert value at before the first node

	newNode := &LinkedList{
		parent: val,
		child:  ll,
	}

	return newNode

}

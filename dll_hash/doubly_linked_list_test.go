package dllhash

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedListHash(t *testing.T) {

	dll := NewDoublyLinkedList()

	int1 := 10
	dll.AddToHead(int1)

	allVal := []int{2, 6, 12}

	for _, val := range allVal {
		dll.AddToTail(val)
	}

	isFound := dll.Find(int1)
	if !isFound {
		t.Errorf("expected to be found, got not found ==> %d", int1)
	}

	int2 := 100
	isFound = dll.Find(int2)
	if isFound {
		t.Errorf("expected to be not found, got found ==> %d", int2)
	}

	allVal = append(allVal, 25)
	for _, val := range allVal {

		if val == 25 {
			isFound = dll.Find(val)
			if isFound {
				t.Errorf("expected to be not found, got found ==> %d", val)
			}

		} else {
			isFound = dll.Find(val)
			if !isFound {
				t.Errorf("expected to be found, got not found ==> %d", val)
			}
		}

	}

	err := dll.AddPrev(2, 109)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = dll.AddPrev(10, 111)
	if err != nil {
		t.Errorf("%v", err)
	}

	dll.TraverseHeadToTail()
	// dll.TraverseTailToHead()

	currNode, err := dll.InfoCurrNode(109) // assume that prev is 2 and next is 10
	if err != nil {
		t.Fatalf("err check info curr node: %v", err)
	}

	if currNode.Next.Val != 2 {
		t.Errorf("next val must be %d, got %d\n", 2, currNode.Next.Val)
	}

	if currNode.Prev.Val != 10 {
		t.Errorf("prev val must be %d, got %d\n", 10, currNode.Prev.Val)
	}

	headNode, err := dll.InfoCurrNode(111) // assume that prev is 2 and next is 10
	if err != nil {
		t.Fatalf("err check info head curr node: %v", err)
	}

	if headNode.Next.Val != 10 {
		t.Errorf("next HEAD must be %d, got %d\n", 10, headNode.Next.Val)
	}

	if headNode.Prev != nil {
		t.Errorf("prev HEAD must be nil, got %v\n", headNode.Prev)
	}

	fmt.Printf("\n pass \n")
}

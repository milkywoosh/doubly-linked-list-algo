package dllhash

import (
	"log"
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

	dll.TraverseToHead()
	// dll.TraverseToTail()

	log.Printf("pass \n")
}

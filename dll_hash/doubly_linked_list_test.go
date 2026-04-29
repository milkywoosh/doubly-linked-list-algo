package dllhash

import (
	"log"
	"testing"
)

func TestDoublyLinkedListHash(t *testing.T) {

	dll := NewDoublyLinkedList()

	int1 := 10
	dll.AddToHead(int1)

	isFound := dll.Find(int1)
	if !isFound {
		t.Errorf("expected to be found, got not found ==> %d", int1)
	}

	int2 := 100
	isFound = dll.Find(int2)
	if isFound {
		t.Errorf("expected to be not found, got found ==> %d", int2)
	}

	log.Printf("pass \n")
}

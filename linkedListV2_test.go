package main

import (
	"log"
	"testing"
)

// func TestConstructLinkedList(t *testing.T) {

// 	var Linky *LinkedList = ConstructLinkedList(2)
// 	if Linky == nil {
// 		t.Errorf("Initiate LinkedList must not be nil")
// 	}

// 	var firstNode int = Linky.current
// 	var expected int = 2
// 	if firstNode != expected {
// 		t.Errorf("got %d, expected %d", firstNode, expected)
// 	}
// }

// func TestAdd(t *testing.T) {

// 	var Linky *LinkedList = ConstructLinkedList(2)

// 	var node1 = Linky
// 	var val = node1.current
// 	if val != 2 {
// 		t.Errorf("node1 must be 2, got %d", val)
// 	}

// 	// add next node
// 	Linky.Add(1)
// 	var node2 = node1.next
// 	var val2 = node2.current
// 	if val2 != 1 {
// 		t.Errorf("node2 must be 1, got %d", val2)
// 	}

// 	// add next after node2
// 	Linky.Add(0)
// 	var node3 = node2.next
// 	var val3 = node3.current
// 	if val3 != 0 {
// 		t.Errorf("val3 must be 0, got %d", val3)
// 	}

// }

// func TestSearch(t *testing.T) {
// 	var Linky *LinkedList = ConstructLinkedList(2)
// 	Linky.Add(10)
// 	Linky.Add(6)

// 	var search0 bool = Linky.Search(0)
// 	if search0 {
// 		t.Errorf("Search(0) must NOT be TRUE, because it never add 0 before, got %t", search0)
// 	}

// 	var search3Fail bool = Linky.Search(3)
// 	t.Logf("LOG -- Search(3) Fail must be FALSE because never adding 3")
// 	if search3Fail == true {
// 		t.Errorf("Search(3) must be FALSE, got %t", search3Fail)
// 	}

// 	Linky.Add(3)
// 	var search3 bool = Linky.Search(3)
// 	t.Logf("LOG -- Search(3) must be TRUE")
// 	if search3 == false {
// 		t.Errorf("Search(3) must be TRUE, got %t", search3)
// 	}
// }

func TestDelete(t *testing.T) {

	var Linky *LinkedList = nil
	// Linky = ConstructLinkedList(5)       // TEST Contrived action
	var errDel1 error = Linky.Delete(10) // currently is nil

	log.Printf("errDel1 ===> %v", errDel1)

	log.Printf("before delete, make sure the data structure is not nil")
	log.Printf("before delete, this test must be error because Linky currently is nil")
	if errDel1 == nil {
		t.Errorf("errDel1 must be error NOT nil, got %v", errDel1)
	}

	Linky = ConstructLinkedList(5)
	log.Printf("after ConstructLinkedList(5)")
	// Linky.Add(10) // TEST Contrived action
	var errDel2 = Linky.Delete(10)
	log.Printf("errDel2 %v", errDel2)

	log.Printf("delete 10 should be failed, and error must NOT be nil, because it never added")
	if errDel2 == nil {
		t.Errorf("failed, because Linky never add 10, and must not be found and cant be deleted")
	}

	Linky.Add(10)
	Linky.Add(8)

	var search8Found bool = Linky.Search(8)
	log.Printf("search 8 found ==> %t", search8Found)
	if search8Found == false {
		t.Errorf("search 8, must be found or TRUE because has been added, got %v", search8Found)
	}

	var errDel3 = Linky.Delete(8)
	log.Printf("delete 8, error must NOT be nil")
	if errDel3 != nil {
		t.Errorf("delete 8, error got: %v, expected <nil> because success search and delete", errDel3)
	}

	// search 8, must be not found
	log.Printf("search 8, must be not found")
	var search8 bool = Linky.Search(8)
	log.Printf("search 8, must be not found %t", search8)
	if search8 == true {
		t.Errorf("search 8, must be not found or FALSE because has been deleted b4")
	}

	log.Printf("search8 after delete: %v ", search8) // expectation nya FALSE/Failed because at line 106 has been deleted

	var search10Found bool = Linky.Search(10)
	if search10Found == false {
		t.Errorf("[FAILS] search 10, must be FOUND or TRUE because has been added b4")
	}

	t.Logf("search10Found : %t ", search10Found) // expectation nya FALSE/Failed because at line 106 has been deleted
}

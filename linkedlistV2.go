package main

import (
	"errors"
	"log"
)

type LinkedList struct {
	current int
	next    *LinkedList
}

func ConstructLinkedList(val int) *LinkedList {
	return &LinkedList{val, nil}
}

func (l *LinkedList) Add(val int) {
	log.Printf("Add %d", val)
	if l != nil {
		if l.next != nil {
			l.next.Add(val)
		} else {
			n := &LinkedList{current: val, next: nil}
			l.next = n
		}
	}
}

func (l *LinkedList) Search(val int) bool {
	// recursive
	// var tmp := &l
	// fmt.Printf("find :%d\n ", val)
	if l != nil {
		log.Printf("look value ==> %d", l.current)

		if l.current == val {
			return true
		} else {
			return l.next.Search(val)
		}
	}
	// last tail which is must be not existed

	return false
}

// list all the value existed
func (l *LinkedList) ShowAll() []int {

	var l_temp *LinkedList = l

	slc := []int{}

	for {
		if l_temp != nil {
			slc = append(slc, l_temp.current)

			// NOTE: perbedaan reassign langsung dan create copy value disini?
			/*
				1. langsung l = l.next
				2. var l_temp = l
					l_temp = l_temp.next
			*/
			l_temp = l_temp.next
		} else {
			return slc
		}
	}

}

func (l *LinkedList) Delete(val int) error {
	log.Printf("Delete %v", val)

	if l != nil {
		// if l.current == val {
		// 	log.Printf("get here => %v", l)
		// 	tmpNext := l.next
		// 	if tmpNext != nil {

		// 		log.Printf("get tmp => %v", tmpNext)
		// 		l.current = tmpNext.current
		// 		l.next = tmpNext.next
		// 		return nil
		// 	} else {
		// 		log.Printf("tmpNext nil")
		// 		return nil
		// 	}
		// } else {
		// }

		log.Printf("lanjut sini 1x, current: %d", l.current)
		if l.next != nil {

			if l.next.current == val {

				log.Printf("wook %v", l.next)
				l.next = nil
				log.Printf("wook1 %v", l.next)
				return nil
			}
		}

		log.Printf("next del recursive")

		return l.next.Delete(val)

	} else {
		return errors.New("failed, current node is nil or not found")
	}

	// note does use Search from existed function will add more resource?
	// var searchVal bool = l.Search(val)
	// if !searchVal {
	// 	errMsg := fmt.Sprintf("failed delete %d, because it doesnt exists", val)
	// 	return errors.New(errMsg)
	// }

	// var tmp *LinkedList = l

}

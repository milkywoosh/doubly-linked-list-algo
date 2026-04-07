package main

import "fmt"

// var l *LinkedList = &LinkedList{current: 10}

/*
func main() {

	l.Add(100)
	l.Add(30)
	l.Add(50)
	l.Add(19)
	l.Add(48)

	l.Print()

	l = l.Prepend(300)
	fmt.Print("limit ============\n")
	fmt.Print("limit ============\n")
	l.Print()
}
*/

func main() {
	// var arg string = "woooo"
	// fmt.Printf("test oke => %s", arg)

	Linky := ConstructLinkedList(1)
	fmt.Println("Linky: ", Linky)
	Linky.Add(11)
	Linky.Add(20)
	Linky.Add(17)

	// fmt.Println(Linky.Search(1))
	// fmt.Println(Linky.Search(11))
	// fmt.Println(Linky.Search(19))

	fmt.Println(Linky.ShowAll())

}

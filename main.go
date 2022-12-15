package main

import (
	ll "ds/linked_list"
	"ds/stack"
	"fmt"
)

func main() {

	// Stack
	s := stack.New(20)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	fmt.Println(s)

	// Linked list
	ll := ll.New()
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(350)
	ll.InsertAtEnd(400)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	fmt.Println(ll)
	ll.DeleteNode(200)
	fmt.Println(ll)
}

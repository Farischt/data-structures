package main

import (
	ll "ds/linked_list"
	"ds/queue"
	"ds/stack"
	"fmt"
)

func main() {
	// Stack
	s := stack.New(20)
	s.Push(10)
	s.Push(20)
	s.Push("hola")
	s.Push(40)
	fmt.Println(s)

	// Queue
	q := queue.New(3)
	q.Enqueue(30)
	q.Enqueue(10)
	q.Enqueue(20)
	item, _ := q.Dequeue()
	fmt.Println(*item)
	fmt.Println(q)

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

package main

import (
	ll "ds/linked_list"
	"ds/queue"
	"ds/stack"
	"ds/tree"
	"fmt"
)

func main() {
	// Stac
	var s stack.IStack = stack.New(20)
	s.Push(10)
	s.Push(20)
	s.Push("hola")
	s.Push(40)
	fmt.Println(s)

	// Queue
	var q queue.IQueue = queue.New(3)
	q.Enqueue(30)
	q.Enqueue(10)
	q.Enqueue(20)
	item, _ := q.Dequeue()
	fmt.Println(*item)
	fmt.Println(q)

	// Linked list
	var ll ll.ILinkedList = ll.New()
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

	// Bst

	var t tree.ITree = tree.New(nil)
	t.Insert(10)
	t.Insert(30)
	t.Insert(50)
	t.Insert(10)
	t.Insert(15)
	r := t.Search(15)
	fmt.Println(r)

	t.InOrderTraversal()
}

package main

import (
	ll "ds/linked_list"
	"ds/queue"
	"ds/stack"
	"ds/tree"
	"fmt"
)

func main() {
	// Stack
	var s stack.IStack[int] = stack.New[int](20)
	s.Push(10)
	s.Push(20)
	s.Push(30)
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
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	r := t.Search(15)
	fmt.Println(r)

	t.InOrderTraversal()
	// 6 is the number of nodes in the tree
	test := t.DepthFirstValues(6)
	fmt.Printf("test: %v", test)
}

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
	var q queue.IQueue[int] = queue.New[int](3)
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
	t.Insert(15)
	t.Insert(20)
	t.Insert(25)
	t.Insert(21)
	t.Insert(22)
	t.Insert(27)
	t.Insert(26)
	r := t.Search(20)
	fmt.Printf("the node searched : %v\n", r)

	t.InOrderTraversal()

	t.Remove(25)
	// 6 is the number of nodes in the tree
	dfv := t.DepthFirstValues(9)
	fmt.Printf("depth first values: %v\n", dfv)

	bfv := t.BreathFirstValues(9)
	fmt.Printf("breath first values: %v\n", bfv)
}

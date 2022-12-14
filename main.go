package main

import (
	ll "ds/linked_list"
	"fmt"
)

func main() {
	ll := ll.New()
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(350)
	ll.InsertAtEnd(400)
	ll.InsertAtEnd(800)
	ll.InsertAtEnd(800)
	ll.InsertAtEnd(900)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(1000)
	ll.InsertAtEnd(200)
	ll.InsertAtEnd(200)
	fmt.Println(ll)
	ll.DeleteNode(200)
	fmt.Println(ll)
}

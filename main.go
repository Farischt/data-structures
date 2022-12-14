package main

import (
	ll "ds/linked_list"
	"fmt"
)

func main() {
	ll := ll.New()
	ll.InsertAtStart(20)
	ll.InsertAtStart(300)
	ll.InsertAtEnd(200)
	fmt.Println(ll)
}

//Reverse a singly linked list in O(n) time.

package main

import (
	"fmt"
	"linkedList"
)

func main() {

	var (
		root  *linkedList.Node = nil
		start *linkedList.Node = nil
	)
	for i := 1; i < 11; i++ {
		start = linkedList.AppendNode(i, start)
		if i == 1 {
			root = start
		}
	}

	start = root

	fmt.Println("The original linked list is: ")
	linkedList.PrintLinkedList(start)
	new_root := reverseLL(root)
	fmt.Println("The reversed linked list is: ")
	linkedList.PrintLinkedList(new_root)
}

func reverseLL(root *linkedList.Node) *linkedList.Node {
	var new_root *linkedList.Node = nil
	for root != nil {
		start := root
		root = root.Next
		start.Next = new_root
		new_root = start
	}
	return new_root
}

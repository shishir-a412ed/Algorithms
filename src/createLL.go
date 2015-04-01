package main

import (
	"fmt"
	"linkedList"
)

func main() {

	var head *linkedList.Node = nil
	head = linkedList.AppendNode(1, head)
	start := head
	head = linkedList.AppendNode(2, head)
	head = linkedList.AppendNode(3, head)
	head = linkedList.AppendNode(4, head)
	head = linkedList.AppendNode(5, head)
	fmt.Println("The linked list is:")
	linkedList.PrintLinkedList(start)
}

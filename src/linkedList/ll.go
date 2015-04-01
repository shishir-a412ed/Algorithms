package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

func createNode(data int, head *Node) *Node {
	head = new(Node)
	head.data = data
	head.next = nil
	return head
}

func AppendNode(data int, head *Node) *Node {
	if head == nil {
		head = createNode(data, head)
	} else {
		head.next = new(Node)
		head = head.next
		head.data = data
		head.next = nil
	}
	return head
}

func PrintLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

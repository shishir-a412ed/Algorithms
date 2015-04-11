package linkedList

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func createNode(data interface{}, head *Node) *Node {
	head = new(Node)
	head.Data = data
	head.Next = nil
	return head
}

func AppendNode(data interface{}, head *Node) *Node {
	if head == nil {
		head = createNode(data, head)
	} else {
		head.Next = new(Node)
		head = head.Next
		head.Data = data
		head.Next = nil
	}
	return head
}

func PrintLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

package main

import (
	"fmt"
	"log"
	"stack"
)

func main() {

	var haystack stack.Stack
	haystack.Push("Hay")
	haystack.Push(-15)
	haystack.Push([]string{"Pin", "Clip", "Needle"})
	haystack.Push(81.52)

	fmt.Printf("The length of the stack is %d\n", haystack.Len())
	item, err := haystack.Top()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Top element of the stack is %v\n", item)
	fmt.Println("Items in the stack are listed below:")
	for {

		item, err = haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

}

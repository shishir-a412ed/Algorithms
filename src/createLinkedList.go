//Create and print a singly linked list of integers.

package main

import (
	"fmt"
	"log"
)

type Link struct {
	data int
	next *Link
}

func main() {

	var (
		num    int
		choice string
	)
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatalln(err)
	}
	head := &Link{num, nil}
	start := head
	for {
		fmt.Println("Do you wish to enter another number {'y' for yes/'n' for no}")
		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			log.Fatalln(err)
		}
		if choice[:1] == "y" {
			fmt.Println("Enter a number")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				log.Fatalln(err)
			}
			head.next = new(Link)
			head = head.next
			head.data = num
			head.next = nil
			continue
		} else {
			break
		}
	}
	fmt.Println("THE LINKED LIST IS: ")

	for start != nil {
		fmt.Printf("%d\n", start.data)
		start = start.next
	}

}

//Test whether a number is even or odd without using % operator.

package main

import (
	"fmt"
	"log"
)

func main() {
	var num int
	fmt.Printf("Enter a number\n")
	if _, err := fmt.Scanf("%d", &num); err != nil {
		log.Fatalln(err)
	}

	if (num/2)*2 == num {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

}

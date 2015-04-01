//Defer calls are pushed onto a stack and retrieved in LIFO order.
//Code taken from tour.golang.org

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Start of loop")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Numbers will be printed after this statement")

}

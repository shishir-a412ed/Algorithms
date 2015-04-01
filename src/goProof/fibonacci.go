//Fibonacci series using go routines and channels.

package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go fibo(ch)
	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}

func fibo(c chan int) {

	i, j := 0, 1
	for {
		c <- j
		i, j = j, i+j
	}
}

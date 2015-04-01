//Split a slice of 12 random integers, and print only those whose value is >500 using goroutines and channels.

package main

import (
	"fmt"
	"math/rand"
)

func checkNum(c []int, ch chan int) {

	for _, x := range c {
		if x > 500 {
			fmt.Printf("%d ", x)
		}
	}
	ch <- 0
}

func main() {

	ch := make(chan int)
	arr := make([]int, 12)
	for i := 0; i < 12; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	fmt.Println(arr)
	go checkNum(arr[:4], ch)
	go checkNum(arr[4:8], ch)
	go checkNum(arr[8:], ch)

	for l := 0; l < 3; l++ {
		<-ch
	}
	close(ch)
}

//Creating 10 worker threads (go routines) and assigning them 50 threads of work, 10 at a time.
//This proof of concept (POC) shows the power of go concurrency.

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func CmdLoad(pause chan int) {

	defer wg.Done()

	for j := range pause {
		fmt.Printf("The value of j is %v ", j)
		fmt.Println("Untar the tar ball into /var/lib/docker/tmp, and then move the image to /var/lib/docker/graph")
		time.Sleep(time.Second * 3)
	}

}

func main() {

	pause := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go CmdLoad(pause)
	}

	for i := 0; i < 50; i++ {
		pause <- i
	}
	close(pause)
	wg.Wait()
	fmt.Println("Batch load done")
}

//Small Proof of concept (POC) to learn how sync.WaitGroup works.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.yahoo.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			fmt.Println(url)
			wg.Done()
		}(url)
		wg.Wait()
	}
}

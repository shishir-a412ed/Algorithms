//Reading an integer using Scanf

package main

import "fmt"

func main() {

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err == nil {
		fmt.Printf("Number entered is %d\n", i)
	}
}

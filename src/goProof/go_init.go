//Small POC to learn how init() works.

package main

import "fmt"

var (
	firstVar  int
	secondVar int
	thirdVar  int
)

func init() {
	firstVar = 0
	secondVar = 1
	thirdVar = 2
	fmt.Printf("first %v, second %v and third %v variables initialized\n", firstVar, secondVar, thirdVar)
}

func main() {
	fmt.Println("Does this execute first or INIT")
	fmt.Printf("Variable accessed in main: %v %v %v\n", firstVar, secondVar, thirdVar)
	fmt.Println("PROGRAM ENDS HERE")
}

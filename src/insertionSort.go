//Algorithm for insertion sort.
//Asymptotic runtime complexity O(n^2).

package main

import (
	"fmt"
	"log"
)

func insertionSort(myList []int) {

	fmt.Println("Input list is:")
	fmt.Println(myList)

	if len(myList) < 1 {
		log.Fatalln("No numbers to be sorted")
	}

	for j := 1; j < len(myList); j++ {
		key := myList[j]
		i := j - 1
		for i >= 0 && myList[i] > key {
			myList[i+1] = myList[i]
			i = i - 1
		}
		myList[i+1] = key
	}

	fmt.Println("Sorted List is:")
	fmt.Println(myList)

}

func main() {

	var (
		myList []int
		choice int
		num    int
	)

	fmt.Println("Enter a list of numbers to be sorted, after each input you ll be prompted with a choice {continue(1)/quit(0)}")

	for {
		fmt.Println("Enter a number")
		if _, err := fmt.Scanf("%d", &num); err != nil {
			log.Fatalln(err)
		}
		myList = append(myList, num)
		fmt.Println("Do you wish to continue or quit")
		if _, err := fmt.Scanf("%d", &choice); err == nil {
			if choice == 1 {
				continue
			} else {
				break
			}
		}
	}

	insertionSort(myList)

}

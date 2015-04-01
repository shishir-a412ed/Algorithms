/*
Reads a number entered on the command line as a string and output the same number onto the console using "big digits".
Please note bigDigits [][]string is not formatted, so in case you want to run the program, you need to format the digits first.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"    000   ",
		"   0	 0  ",
		"  0     0 ",
		"  0     0 ",
		"  0     0 ",
		"   0   0  ",
		"    000   "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111 "},
	{" 222 ", "2   2", "  2 ", "  2 ", " 2 ", "2   ", "22222"},
	{"3", "3", "3", "3", "3", "3", "3"},
	{"4", "4", "4", "4", "4", "4", "4"},
	{"5", "5", "5", "5", "5", "5", "5"},
	{"6", "6", "6", "6", "6", "6", "6"},
	{"7", "7", "7", "7", "7", "7", "7"},
	{"8", "8", "8", "8", "8", "8", "8"},
	{" 9999", "9   9", "9   9", " 9999", "   9", "   9", "   9"},
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if digit >= 0 && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("Invalid whole number")
			}
		}
		fmt.Println(line)
	}
}

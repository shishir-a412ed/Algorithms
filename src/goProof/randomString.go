//Program to generate a random string of length n

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(length int) string {

	source := rand.NewSource(int64(time.Now().Nanosecond()))
	rd := rand.New(source)
	b := make([]rune, length)
	for i := 0; i < length; i++ {
		b[i] = letters[rd.Intn(len(letters))]
	}
	return string(b)
}
func main() {
	fmt.Println(randString(7))
}

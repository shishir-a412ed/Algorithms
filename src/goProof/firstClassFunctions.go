// A very simple example to learn first class functions in go.
// code taken from http://ernestmicklei.com/2012/11/13/first-class-functions-in-go/

package main

import "fmt"

func CallWith(f func(string), who string) {
	f(who)
}

type FunctionHolder struct {
	Function func(string)
}

func main() {
	holder := &FunctionHolder{func(who string) {
		fmt.Println("Hello,", who)
	}}
	CallWith(holder.Function, "World")
}

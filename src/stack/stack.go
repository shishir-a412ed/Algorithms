package stack

import (
	"fmt"
	"errors"
	)

type Stack []interface{}

func (stack Stack) Len() int{
return len(stack)
}

func (stack Stack) Top() (interface{},error){
if len(stack)==0{
	return nil,errors.New("Cannot Top() an empty stack")
}
return stack[len(stack)-1],nil
}

func (stack *Stack) Push(x interface{}){
*stack = append(*stack,x)
}

func (stack *Stack) Pop() (interface{},error){

if len(*stack)==0{
	return nil,errors.New("Cannot Pop() an empty stack")
}
x=*stack[len(*stack)-1]
*stack = *stack[:len(*stack)-1]
return x,nil
}

package main 

import (
	"fmt"
	"strconv"
)

const stackLenth = 4

type Stack [stackLenth]string

var stackLastKey int = 0

func (self Stack) String() (stackInfo string) {

	for key, value := range self {
		stackInfo += "[" + strconv.Itoa(key) + ":" + value + "]"
	}

	return
}

func (self *Stack) Push(s string) {
	if (stackLastKey >= stackLenth) {
		fmt.Println("index out of range")
	}
	self[stackLastKey] = s
	stackLastKey ++
}

func (self *Stack) Pop() {
	if (stackLastKey <= 0) {
		fmt.Println("index out of range")
	}
	self[stackLastKey-1] = ""
	stackLastKey --
}

func main() {
	var stack Stack

	stack.Push("a")

	fmt.Println(stack)

	stack.Push("b")
	stack.Push("c")

	fmt.Println(stack)

	stack.Pop()

	fmt.Println(stack)
}
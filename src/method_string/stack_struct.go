package main 

import (
	"fmt"
	"strconv"
)

const stackLenth = 4

type Stack struct {
	lastKey int
	stack [stackLenth]string
}

func (self Stack) String() (stackInfo string) {
// func (self *Stack) String() (stackInfo string) {  这样就没办法执行到

	for key, value := range self.stack {
		stackInfo += "[" + strconv.Itoa(key) + ":" + value + "]"
	}

	return
}

func (self *Stack) Push(s string) {
	if (self.lastKey >= stackLenth) {
		fmt.Println("index out of range")
	}
	self.stack[self.lastKey] = s
	self.lastKey ++
}

func (self *Stack) Pop() {
	if (self.lastKey <= 0) {
		fmt.Println("index out of range")
	}
	self.stack[self.lastKey-1] = ""
	self.lastKey --
}

func main() {
	var stack Stack

	stack.lastKey = 0

	stack.Push("a")

	fmt.Println(stack)

	stack.Push("b")
	stack.Push("c")

	fmt.Println(stack)

	stack.Pop()

	fmt.Println(stack)
}
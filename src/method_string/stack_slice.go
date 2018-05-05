package main 

import (
	"fmt"
	"strconv"
)

const stackLenth = 4

type Stack []string

func (self Stack) Size() (lenth int) {
	lenth = len(self)
	return
}

func (self Stack) String() (stackInfo string) {

	for key, value := range self {
		stackInfo += "[" + strconv.Itoa(key) + ":" + value + "]"
	}

	return
}

// func (self Stack) Push(s string) {
// 	//这里切片内容发生变化没有影响到本身的值，问题还在调查中。。。
// 	self = append(self, s)
// }

func (self Stack) Pop() {
	lastKey := self.Size()
	if (lastKey > 0) {
		self = self[:lastKey]
	}
	
}

func main() {
	var stack  = make(Stack, 0, stackLenth)

	stack.Push("a")

	fmt.Println(stack)

	stack.Push("b")

	fmt.Println(stack)

	stack.Pop()

	fmt.Println(stack)


}
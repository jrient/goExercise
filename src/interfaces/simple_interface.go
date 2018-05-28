package main 

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	value int
}

func (self Simple) Get() int {
	return self.value
}

func (self *Simple) Set(value int) {
	self.value = value
}

func test_simple(s Simpler) {
	value := s.Get()
	fmt.Printf("before set, value = %d \n", value)
	s.Set(10)
	value = s.Get()
	fmt.Printf("after set, value = %d \n", value)
}

func main() {
	s := &Simple{1}
	test_simple(s)
}
package main 

import (
	"fmt"
)

type Day int

var weeks = []string{"Sunday", "Monday", "Tuesday", "Wendnesday", "Friday", "Saturday"}

const (
	SU = iota
	MO
	TU
	WE
	FR
	SA
)

func (self Day) String() string {
	return weeks[self]
}

func main() {
	var day Day = 2
	fmt.Println(day)
	fmt.Println(TU)
}
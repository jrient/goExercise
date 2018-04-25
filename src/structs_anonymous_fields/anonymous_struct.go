package main

import (
	"fmt"
)

type Stru struct {
	a float32
	int
	string
}

func main() {
	stru2 := Stru{2.5, 1, "hello"}

	fmt.Println(stru2)
}

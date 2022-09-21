package main

import "fmt"

func main() {
	fv := func(){
		fmt.Println("hello world")
	}

	fv()
	fmt.Printf("fv is of type %T and has value %v", fv, fv)
}

/*

hello world
fv is of type func() and has value 0x4826e0

*/
package main

import (
    "fmt"
)


func main() {
    out := make(chan int)
    // out <- 2
	go dump(out)
	fmt.Println(<-out)
}

func dump(ch chan int) {
	ch <- 2
}

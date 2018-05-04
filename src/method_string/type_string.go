package main

import (
    "fmt"
    "strconv"
)

type T struct {
    a int
    b float32
    c string
}

func (this *T) String() string {
	return strconv.Itoa(this.a) + " / " + 
	strconv.FormatFloat(float64(this.b), 'f', 6, 32) + " / " + 
	strconv.Quote(this.c)
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t)
}
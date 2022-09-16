package main

import (
    "fmt"
    "strconv"
)

func main() {
    var orig string = "ABC"
	// var orig string = "123"
    // var an int
    var newS string
    // var err error

    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
    an, err := strconv.Atoi(orig)
    if err != nil {
        fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
        return
    } 
    fmt.Printf("The integer is %d\n", an)
    an = an + 5
    newS = strconv.Itoa(an)
    fmt.Printf("The new string is: %s\n", newS)
}

/*

he size of ints is: 64
orig ABC is not an integer - exiting with error

*/
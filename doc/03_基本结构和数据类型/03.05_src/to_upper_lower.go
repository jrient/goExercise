package main

import (
    "fmt"
    "strings"
)

func main() {
    var orig string = "Hey, how are you George?"
    var lower string
    var upper string

    fmt.Printf("The original string is: %s\n", orig)
    lower = strings.ToLower(orig)
    fmt.Printf("The lowercase string is: %s\n", lower)
    upper = strings.ToUpper(orig)
    fmt.Printf("The uppercase string is: %s\n", upper)
}

/*
The original string is: Hey, how are you George?
The lowercase string is: hey, how are you george?
The uppercase string is: HEY, HOW ARE YOU GEORGE?
*/
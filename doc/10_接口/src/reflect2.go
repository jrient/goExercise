package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    // setting a value:
    // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
    fmt.Println("settability of v:", v.CanSet())
    v = reflect.ValueOf(&x) // Note: take the address of x.
    fmt.Println("type of v:", v.Type())
    fmt.Println("settability of v:", v.CanSet())
    v = v.Elem()
    fmt.Println("type of v:", v.Type())
    fmt.Println("The Elem of v is: ", v)
    fmt.Println("settability of v:", v.CanSet())
    v.SetFloat(3.1415) // this works!
    fmt.Println(v.Interface())
    fmt.Println(v)
}

/*
settability of v: false
type of v: *float64
settability of v: false
type of v: float64
The Elem of v is:  3.4
settability of v: true
3.1415
3.1415
*/
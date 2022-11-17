package main

import (
    "fmt"
    "math"
)

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    var areaIntf Shaper
    sq1 := new(Square)
    sq1.side = 5

    areaIntf = sq1

    fmt.Printf("The sq1 type is: %T\n", sq1)
    fmt.Printf("The areaIntf type is: %T\n", areaIntf)
    
    switch t := areaIntf.(type) {
    case *Square:
        fmt.Printf("Type Square %T with value %v\n", t, t)
    case *Circle:
        fmt.Printf("Type Circle %T with value %v\n", t, t)
    case nil:
        fmt.Printf("nil value: nothing to check?\n")
    default:
        fmt.Printf("Unexpected type %T\n", t)
    }
    
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
    return ci.radius * ci.radius * math.Pi
}

/*
Type Square *main.Square with value &{5}
*/
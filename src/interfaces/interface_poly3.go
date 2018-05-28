package main

import (
    "fmt"
    "math"
)

type Shaper interface {
	Abs2(float64, float64) float64
}

type Shape struct{

}

func (self Shape) Abs2(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

type Point struct {
    x, y float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
    Point
    Shape
    name string
}

func main() {
    n := &NamedPoint{Point{3, 4},Shape{},  "Pythagoras"}
    fmt.Println(n.Abs()) // 打印5
    fmt.Println(n.Abs2(n.x, n.y)) // 打印5
}
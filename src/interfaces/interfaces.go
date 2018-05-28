package main

import "fmt"
import "github.com/liudng/godump"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func main() {
    sq1 := new(Square)
    sq1.side = 5

    var areaIntf Shaper
    areaIntf = sq1
    //定义接口变量，此变量type=nil 当Square变量赋值给接口变量后，areaIntf的type也变成了Square
    //如果Square没有被判定继承Shaper,则会提示 cannot convert sq1 (type *Square) to type Shaper
    // shorter,without separate declaration:
    // areaIntf := Shaper(sq1)
    // or even:
    // areaIntf := sq1
    fmt.Printf("The square has area: %f\n", areaIntf.Area())
    godump.Dump(areaIntf)

    fmt.Printf("%T",areaIntf)
}
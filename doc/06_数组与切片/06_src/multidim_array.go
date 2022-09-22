package main

import "fmt"

const (
    WIDTH  = 100
    HEIGHT = 50
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++ {
            screen[x][y] = 0
        }
    }

	fmt.Println(screen)
}

package main 

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (this Celsius) String() string {
	return strconv.Itoa(int(this)) + "°C"
}

func main() {
	var cel Celsius
	cel = 12.3

	fmt.Println(cel)
}
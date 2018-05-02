package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car
}

func (this *Car) numberOfWheels() (wheels int) {
	wheels = this.wheelCount
	return
}

func (this *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi merkel")
}

func main() {
	mercedes := new(Mercedes)
	mercedes.wheelCount = 4

	// mercedes2 := Mercedes{Car{4, Engine}}
	// 不知道怎么初始化

	fmt.Println(mercedes.numberOfWheels())
	mercedes.sayHiToMerkel()
	// fmt.Println(mercedes2.numberOfWheels())
}

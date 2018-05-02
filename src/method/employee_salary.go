package main

import (
	"fmt"
)

type Employee struct {
	salary int
}

func (this Employee) giveRaise(raise float32) (salary int) {
	salary = this.salary + int(float32(this.salary)*raise)
	return
}

func main() {
	emp := Employee{1000}
	fmt.Println(emp.giveRaise(0.2))
}

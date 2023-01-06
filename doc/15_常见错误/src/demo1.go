package main

import (
	"fmt"
)

func main() {
	var remeber bool = false
	var something bool = true
	if something {
		// 短声明导致 remeber是一个局部变量
		// remeber := true

		// 正确写法
		remeber = true
		fmt.Println(remeber)
	}
	fmt.Println(remeber)
}

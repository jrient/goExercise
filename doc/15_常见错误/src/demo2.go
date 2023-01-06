package main

import (
	"fmt"
)

func main() {
	fmt.Println(shadow())
}

func shadow() (result string) {
	// x.result := f1()
	x, res1 := f1()

	if res1 == "bb" {
		fmt.Println(x)
		result = res1
		return
	}

	// if y, result := f2(); result != "cc" {
	if y, res2 := f2(); res2 == "cc" {
		result = res2
		return
	} else {
		fmt.Println(y)
	}
	return
}

func f1() (string, string) {
	return "aa", "bb"
}

func f2() (string, string) {
	return "cc", "dd"
}

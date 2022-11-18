package main

import (
	"fmt"
	"reflect"
)

type MyInt int

var m MyInt = 5

func main() {
	v := reflect.ValueOf(m)

	fmt.Println(v.Kind())

	fmt.Println(v.Interface())
}

/*
int
5
*/
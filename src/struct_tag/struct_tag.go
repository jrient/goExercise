package main

import (
	"fmt"
	// "github.com/liudng/godump"
	// "os"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
		// os.Exit(1)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)

	// godump.Dump(ttType)

	ixField := ttType.Field(ix)

	// godump.Dump(ixField)

	fmt.Printf("%v\n", ixField.Tag)
}

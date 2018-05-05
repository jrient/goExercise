package main 

import (
	"fmt"
)

type TZ string

var timezonesMap = map[string]string{
	"ACDT": "Australian Central Daylight Time",
	"ACT": "Acre Time",
	"UTC": "Universal Greenwich time"}

func (self TZ) String() string {
	return timezonesMap[string(self)]
}

func main() {
	var tz TZ = "ACT"

	fmt.Println(tz)
}

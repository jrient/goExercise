// gob2.go
package main

import (
    "encoding/gob"
    "log"
    "os"
	"bufio"
	"fmt"
)

type Address struct {
    Type             string
    City             string
    Country          string
}

type VCard struct {
    FirstName   string
    LastName    string
    Addresses   []*Address
    Remark      string
}

var content string

func main() {
    pa := &Address{"private", "Aartselaar","Belgium"}
    wa := &Address{"work", "Boom", "Belgium"}
    vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}
    // fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
    // using an encoder:
    file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
    defer file.Close()
    enc := gob.NewEncoder(file)
    err := enc.Encode(vc)
    if err != nil {
        log.Println("Error in encoding gob")
    }

	//read
	file2, err := os.Open("vcard.gob")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	reader := bufio.NewReader(file2)
	var decStr VCard
	dec := gob.NewDecoder(reader)
	err2 := dec.Decode(&decStr)
	fmt.Println(decStr)
	fmt.Println(decStr.Addresses[0],decStr.Addresses[1])
	if err2 != nil {
		log.Println("Error in decode gob")
	}
}

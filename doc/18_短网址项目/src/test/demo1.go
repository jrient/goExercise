package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
)

type record struct {
	Key, Value string
}

func main() {
	f, err := os.OpenFile("./store.gob", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("open file failed", err)
	}

	if err := save(f); err != nil {
		log.Fatal("save file failed", err)
	}

	if err := load(f); err != nil {
		log.Fatal("load file failed", err)
	}

}

func save(f *os.File) error {
	e := gob.NewEncoder(f)
	return e.Encode(record{"a", "c"})
}

func load(f *os.File) error {
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	d := gob.NewDecoder(f)
	var err error
	for err == nil {
		end := read(f, d)
		if end == true {
			return nil
		}
	}
	return err
}

func read(f *os.File, d *gob.Decoder) bool {
	var r record
	if err := d.Decode(&r); err == nil {
		fmt.Println("load recode:", r)
	} else {
		if err == io.EOF {
			return true
		} else {
			// log.Fatal("load file error:", err)
		}
	}
	return false
}

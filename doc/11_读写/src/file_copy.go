// filecopy.go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    CopyFile("products2.txt", "products2_copy.txt")
    fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
		fmt.Println(err)
        return
    }
    defer src.Close()

    dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
		fmt.Println(err)
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

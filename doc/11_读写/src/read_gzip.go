package main

import (
    "fmt"
    "bufio"
    "os"
    "compress/gzip"
)

func main() {
    fName := "products2.txt.gz"
    // fName := "products2.txt"
    var r *bufio.Reader
    fi, err := os.Open(fName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
            err)
        os.Exit(1)
    }
    fz, err := gzip.NewReader(fi)
    if err != nil {
		fmt.Println("do not reading in gz:", err, fi, fz)
        r = bufio.NewReader(fi)
    } else {
        r = bufio.NewReader(fz)
    }

    for {
        line, err := r.ReadString('\n')
        fmt.Println(line)
        if err != nil {
            fmt.Println("Done reading file")
            os.Exit(0)
        }
    }
}

/*
姓名 小红 小兰

年龄 18 19

性别 男 女
Done reading file
*/
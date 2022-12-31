package main

import (
    "flag"
    "fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

// last 相当于上一个的意思
func f(last, current chan int) {
    temp := 1 + <-current // 当 current 没有值时,这里会被阻塞
    last <- temp
}
func main() {
    first := make(chan int)

    last := first
    for i := 0; i < *ngoroutine; i++ {
        current := make(chan int)

        // 将上一次循环创建的 chan,和本次循环的 chan 一起交给函数, 函数会帮我们完成 last <- 1+ <- curr 的过程
        go f(last, current)

        // 记录本次循环中的 right,给下一次循环创建使用
        last = current
    }

    // 开始链接
    last <- 0

    x := <-first // wait for completion 等待完成

    fmt.Println(x)
    // 结果： 100000 ， 大约 1,5s （我实际测试只用了不到200ms）
}

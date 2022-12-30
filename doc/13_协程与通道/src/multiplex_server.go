package main

import "fmt"

type Request struct {
    a, b   int
    replyc chan int // 请求中的回复频道

}

type binOp func(a, b int) int

func run(op binOp, req *Request) {

    req.replyc <- op(req.a, req.b)

}

func server(op binOp, service chan *Request) {

    for {

        req := <-service // 请求到达这里

        // 开启请求的 Goroutine ：

        go run(op, req) // 不要等待 op

    }

}

func startServer(op binOp) chan *Request {

    reqChan := make(chan *Request)

    go server(op, reqChan)

    return reqChan

}

func main() {

    adder := startServer(func(a, b int) int { return a + b })

    const N = 100

    var reqs [N]Request

    for i := 0; i < N; i++ {

        req := &reqs[i]

        req.a = i

        req.b = i + N

        req.replyc = make(chan int)

        adder <- req

    }

    // 校验：

    for i := N - 1; i >= 0; i-- { // 顺序无所谓

        if <-reqs[i].replyc != N+2*i {

            fmt.Println("fail at", i)

        } else {

            fmt.Println("Request ", i, "is ok!")

        }

    }

    fmt.Println("done")

}


/*
Request  99 is ok!
Request  98 is ok!
Request  97 is ok!
Request  96 is ok!
Request  95 is ok!
Request  94 is ok!
Request  93 is ok!
Request  92 is ok!
Request  91 is ok!
Request  90 is ok!
Request  89 is ok!
Request  88 is ok!
Request  87 is ok!
Request  86 is ok!
Request  85 is ok!
Request  84 is ok!
Request  83 is ok!
Request  82 is ok!
Request  81 is ok!
Request  80 is ok!
Request  79 is ok!
Request  78 is ok!
Request  77 is ok!
Request  76 is ok!
Request  75 is ok!
Request  74 is ok!
Request  73 is ok!
Request  72 is ok!
Request  71 is ok!
Request  70 is ok!
Request  69 is ok!
Request  68 is ok!
Request  67 is ok!
Request  66 is ok!
Request  65 is ok!
Request  64 is ok!
Request  63 is ok!
Request  62 is ok!
Request  61 is ok!
Request  60 is ok!
Request  59 is ok!
Request  58 is ok!
Request  57 is ok!
Request  56 is ok!
Request  55 is ok!
Request  54 is ok!
Request  53 is ok!
Request  52 is ok!
Request  51 is ok!
Request  50 is ok!
Request  49 is ok!
Request  48 is ok!
Request  47 is ok!
Request  46 is ok!
Request  45 is ok!
Request  44 is ok!
Request  43 is ok!
Request  42 is ok!
Request  41 is ok!
Request  40 is ok!
Request  39 is ok!
Request  38 is ok!
Request  37 is ok!
Request  36 is ok!
Request  35 is ok!
Request  34 is ok!
Request  33 is ok!
Request  32 is ok!
Request  31 is ok!
Request  30 is ok!
Request  29 is ok!
Request  28 is ok!
Request  27 is ok!
Request  26 is ok!
Request  25 is ok!
Request  24 is ok!
Request  23 is ok!
Request  22 is ok!
Request  21 is ok!
Request  20 is ok!
Request  19 is ok!
Request  18 is ok!
Request  17 is ok!
Request  16 is ok!
Request  15 is ok!
Request  14 is ok!
Request  13 is ok!
Request  12 is ok!
Request  11 is ok!
Request  10 is ok!
Request  9 is ok!
Request  8 is ok!
Request  7 is ok!
Request  6 is ok!
Request  5 is ok!
Request  4 is ok!
Request  3 is ok!
Request  2 is ok!
Request  1 is ok!
Request  0 is ok!
done
*/
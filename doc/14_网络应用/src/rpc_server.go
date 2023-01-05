// rpc_server.go
//  原文注释已经被删除，因为和此代码没有关系，个人猜测是作者在这个示例修改之前的代码。

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"./rpc_objects"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}

/* 输出:
启动程序 E:/Go/GoBoek/code_examples/chapter_14/rpc_server.exe ...

** after 5 s: **
End Process exit status 0
*/

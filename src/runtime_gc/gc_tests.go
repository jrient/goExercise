package main 

import (
	"runtime"
	"fmt"
)


func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	fmt.Printf("%d Kb\n", m.TotalAlloc / 1024)
	fmt.Printf("%d Kb\n", m.Frees / 1024)
	fmt.Printf("%d Kb\n", m.HeapIdle / 1024)


	//如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
	//runtime.SetFinalizer(obj, func(obj *typeObj))
}
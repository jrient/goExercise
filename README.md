# Golang 练习记录

标签（空格分隔）： go

----

书籍 《[Go入门指南》][1]

- 10.3 [使用自定义包中的结构体][2]
    - struct_pack
        - 每个包必须建立一个独立的目录
        - 包名与目录名必须相同
- 10.4 [带标签的结构体][3]
    - struct_tag
        - 简单的断点调试 os.Exit(1)
        - tag 加在 type 之后
- 10.5 [匿名字段和内嵌结构体][4]
    - structs_anonymous_fields
        - 注意继承的结构体字段命名
        - 匿名字段用{structName}.{typeName} 使用
- 10.6 [方法][5]
    - method
        - 不知道如何使用 ':=' 来初始化一个带接口的结构体
        - 在类型中嵌入功能可以使用'聚合'和'内嵌',定义区别在于包含的字段是否具备名称；使用的区别在于调用的方法。
- 10.7 [类型的 String() 方法和格式化描述符][6]
    - method_sting
        - 用于定义标准格式化输出，适用于fmt.Print() fmt.Println()默认描述符和描述符%v
        - 不要在 String() 方法里面调用涉及 String() 方法的方法，它会导致意料之外的错误
        - [strconv包的使用][7]
        - 关于代码过长[续行][8]
        - func (self *Stack) String() string { return "a" }  使用 *Stack 执行输出就走不到这里
- 10.8 [垃圾回收和 SetFinalizer][9]
    - runtime_gc(实际上只是一些runtime包功能的测试)
        - [runtime文档][10]
        - [runtime中文文档][11]
        - [runtime中文详解][12]
- 11.1 [接口是什么][13]
    - interfaces
        - 按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成
        - 接口变量的类型随着赋值变量的类型而改变（前提是赋值变量实现了接口）
- 11.3 [断言类型][14]
    - type_interfaces
        - 如果忽略 interfaceValue.(*Type) 中的 * 号，会导致编译错误
- 11.4 [type_switch类型判断](https://zengweigang.gitbooks.io/core-go/content/eBook/11.4.html)
    - type_switch_interfaces
        - type-switch 不允许有 fallthrough
- 11.5 [使用方法集与接口](https://zengweigang.gitbooks.io/core-go/content/eBook/11.6.html)
    - 指针方法可以通过指针调用
    - 值方法可以通过值调用
    - 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
    - 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
- 11.5 [第一个例子：使用 Sorter 接口排序](https://zengweigang.gitbooks.io/core-go/content/eBook/11.7.html)
    - interfaces_sort_t

        
  [1]: https://legacy.gitbook.com/book/zengweigang/core-go/details
  [2]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.3.html
  [3]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.4.html
  [4]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.5.html
  [5]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.6.html
  [6]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.7.html
  [7]: http://www.cnblogs.com/golove/p/3262925.html
  [8]: https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/
  [9]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.8.html
  [10]: https://golang.org/pkg/runtime/#MemStatsType
  [11]: https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/108.html
  [12]: https://zhuanlan.zhihu.com/p/27328476
  [13]: https://zengweigang.gitbooks.io/core-go/content/eBook/11.1.html
  [14]: https://zengweigang.gitbooks.io/core-go/content/eBook/11.3.html

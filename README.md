#Golang 练习记录

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


  [1]: https://legacy.gitbook.com/book/zengweigang/core-go/details
  [2]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.3.html
  [3]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.4.html
  [4]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.5.html
  [5]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.6.html
  [6]: https://zengweigang.gitbooks.io/core-go/content/eBook/10.7.html
  [7]: http://www.cnblogs.com/golove/p/3262925.html
  [8]: https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/
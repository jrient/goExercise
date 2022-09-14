## IDE和工具

### IDE介绍

[这里介绍了Go开发环境的基本要求](https://learnku.com/docs/the-way-to-go/basic-requirements-for-go-development-environment/3573)，有兴趣的可以了解下。

[这里介绍了一些大牛们推荐的IDE](https://learnku.com/docs/the-way-to-go/editors-and-integrated-development-environments/3574)，可以选择自己喜欢的使用。

我这边使用的是万能的**VSCode**，安装上**Go**插件即可，具体的功能可以自己查看文档。

### 代码调试

目前主要的调试器是`gdb`,更原始的方式是使用**变量打印**进行调试。

以下是可用的打印命令
- print
- println
- fmt.Print
- fmt.Println
- fmt.printf

另外(还没学，后面补上详细内容)：
- 可以使用 panic 语句来获取栈跟踪信息（直到 panic 时所有被调用函数的列表）
- 可以使用关键字 defer 来跟踪代码执行过程

### 构建和运行
- `go build` 编译并安装自身包和依赖包
- `go install` 安装自身包和依赖包


### 格式化工具

[gofmt官方文档](https://pkg.go.dev/cmd/gofmt)

总的来说，就是go官方不想要大家有太多花里胡哨的风格，就设置了一个官方的代码格式化工具`gofmt`，并且建议在每次编译之前格式化你的代码。

使用格式化命令，对目录下的所有.go文件进行格式化并**覆盖原文件**:
> gofmt -w *.go

gofmt也提供类似于`sed -i`的功能：

```shell gofmt -r '原始内容 -> 替换内容'  -w replace.go ```

例如下面的代码会将`replace.go`文件中的所有`abc`替换为`123`:

```shell gofmt -r 'abc -> 123' -w replace.go```

### 生成代码文档

[go doc 官方文档](https://golang.org/cmd/godoc/)

`go doc` 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。

它也可以作为一个提供在线文档浏览的 web 服务器，golang.org 就是通过这种形式实现的

- `go doc package` 获取包的文档注释，例如：go doc fmt 会显示使用 godoc 生成的 fmt 包的文档注释。
- `go doc package/subpackage` 获取子包的文档注释，例如：go doc container/list。
- `go doc package function` 获取某个函数在某个包中的文档注释，例如：go doc fmt Printf 会显示有关 fmt.Printf() 的使用说明。
- `godoc -http=:6060` 可以打开一个本地文档浏览web服务

### 其他工具

- `go install` 是安装 Go 包的工具，类似 Ruby 中的 rubygems。主要用于安装非标准库的包文件，将源代码编译成对象文件。
- `go fix` 用于将你的 Go 代码从旧的发行版迁移到最新的发行版，它主要负责简单的、重复的、枯燥无味的修改工作，如果像 API 等复杂的函数修改，工具则会给出文件名和代码行数的提示以便让开发人员快速定位并升级代码。Go 开发团队一般也使用这个工具升级 Go 内置工具以及 谷歌内部项目的代码。go fix 之所以能够正常工作是因为 Go 在标准库就提供生成抽象语法树和通过抽象语法树对代码进行还原的功能。该工具会尝试更新当前目录下的所有 Go 源文件，并在完成代码更新后在控制台输出相关的文件名称。
- `go test` 是一个轻量级的单元测试框架。


### 性能说明
[Go 性能说明](https://learnku.com/docs/the-way-to-go/go-performance-description/3580)

### 其他语言交互
[Go 与其他语言进行交互](https://learnku.com/docs/the-way-to-go/interact-with-other-languages/3581)
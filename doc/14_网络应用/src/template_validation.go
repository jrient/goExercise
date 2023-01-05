package main

import (

    "text/template"

    "fmt"

)

func main() {

    tOk := template.New("ok")

    // 一个有效的模板，所以 Must 时候不会出现恐慌（panic）

    template.Must(tOk.Parse("/*这是一个注释 */ some static text: {{ .Name }}"))

    fmt.Println("The first one parsed OK.")

    fmt.Println("The next one ought to fail.")

    tErr := template.New("error_template")

    template.Must(tErr.Parse("some static text {{ .Name }"))

}


/*
The first one parsed OK.
The next one ought to fail.
panic: template: error_template:1: unexpected "}" in operand

goroutine 1 [running]:
text/template.Must(...)
        /usr/local/go/src/text/template/helper.go:26
main.main()
        template_validation.go:25 +0x2aa
exit status 2
*/
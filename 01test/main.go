package main

import "fmt"

// 推荐使用小驼峰命名法

var (
	name string
	age  int
	isOK bool
)

func main() {
	name = "理想"
	age = 16
	isOK = true
	fmt.Printf("name:%s", name) // 格式化输出
	fmt.Println(age)            // 打印完毕最后输出一个空行(加了一个\n)
	fmt.Println()               // 打印一个空行
	fmt.Print(isOK)             // 正常打印
}

// [Running] go run "c:\code_all\goCode_all\01test\main.go"
// name:理想16

// true
// [Done] exited with code=0 in 1.403 seconds

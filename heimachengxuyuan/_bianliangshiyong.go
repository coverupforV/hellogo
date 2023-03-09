package main

import "fmt"

func main() {
	//变量的声明
	//1.声明格式   var 变量名 类型名
	//只有声明没有初始化的的变量初始值为0
	//同一个{}里声明的变量名唯一
	var a int
	//声明的变量必须使用，不然会报错
	fmt.Println("a=", a)

	//可以声明多个变量
	//var b, c int
	a = 10
	fmt.Println("a=", a)

}

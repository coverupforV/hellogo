package main

import "fmt"

func main() {
	//变量的声明
	//1.声明格式   var 变量名 类型名
	//只有声明没有初始化的的变量初始值为0
	//同一个{}里声明的变量名唯一
	var a int
	//声明的变量必须使用，不然会报错
	fmt.Println("a =", a)

	//可以声明多个变量
	//var b, c int
	a = 10 //赋值
	fmt.Println("a =", a)

	var b int = 10 //2.声明的时候就赋值
	b = 20
	fmt.Println("b = ", b)

	//3。自动推导类型，必须初始化，由初始化的值确定类型（常用）
	c := 30
	//%T输出变量类型
	fmt.Printf("c type is %T\n", c)

}

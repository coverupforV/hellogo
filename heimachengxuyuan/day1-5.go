package main

import "fmt"

func main() {

	//常量：程序运行期间不可以改变的量，用const声明

	const a int = 1
	//a = 20  error 常量不允许修改
	fmt.Println("a = ", a)

	const b = 11.2 //没有使用：=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}

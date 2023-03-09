package main

import "fmt"

func main() {
	//这样要先声明再赋值
	var a int
	a = 10
	fmt.Println("a = ", a)
	//可以用：=直接赋值，会自动推导类型，同一个变量名只能用一次
	b := 10
	fmt.Println("b = ", b)

	//b：=30  这样不行，前边有变量b，不能再使用b了

}

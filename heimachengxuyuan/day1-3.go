package main

import "fmt"

func main() {
	//println与printf区别
	//Println 一段一段处理,按顺序依次输出，自动加换行
	a := 10
	fmt.Println("a = ", a)
	//printf是格式化输出，把a的内容放在%d的位置
	fmt.Printf("a = %d", a)

}

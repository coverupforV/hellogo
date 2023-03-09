package main

import "fmt"

func main() {
	//1.iota常量自动生成器，每隔一行，自动加一
	//2.iota给常量赋值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)

	//iota遇到const，重置为0
	const d = iota //0
	fmt.Printf("d=%d\n", d)
}

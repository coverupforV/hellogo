package main

import "fmt"

func main() {
	//字符类型
	var ch byte
	ch = 97
	fmt.Printf("%c", ch)

	//字符，单引号
	ch = 'a' //也可

	//字符串类型,也可以使用自动推导类型
	var str1 string
	str1 = "123"
	fmt.Println(str1)
	//字符串长度len（）
	fmt.Println(len(str1))
}

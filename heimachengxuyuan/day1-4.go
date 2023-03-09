package main

import (
	"fmt"
)

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//多重赋值
	a, b := 10, 20

	//交换两个变量的值
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)

	//匿名变量_ （下划线）
	i, j := 10, 11
	var temp int
	temp, _ = i, j
	fmt.Println("temp = ", temp)
	//匿名变量丢弃数据不处理，配合函数返回值使用有优势

	//匿名变量举例
	var c, d, e int
	c, d, e = test()
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	_, d, _ = test()
	fmt.Printf("d = %d\n", d)
}

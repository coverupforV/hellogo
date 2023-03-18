package main

import (
	"fmt"
	//"time"
	"runtime"
)

// func newTask() {
// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("new goroutine: i = %d\n", i)
// 		time.Sleep(1 * time.Second) //延时1s
// 	}
// }

// func main() {
// 	//创建一个 goroutine，启动另外一个任务
// 	go newTask()
// 	i := 0
// 	//main goroutine 循环打印
// 	for {
// 		i++
// 		fmt.Printf("main goroutine: i = %d\n", i)
// 		time.Sleep(1 * time.Second) //延时1s
// 	}
// }

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 终止当前 goroutine, import "runtime"
			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}() //不要忘记()

	//死循环，目的不让主goroutine结束
	for {
	}
}

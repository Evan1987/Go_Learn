package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)  // 如果数量大于1则为并行，否则为排队并发

	// wg 用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已完成
		// 保证每个goroutine一旦完成工作就调用Done
		defer wg.Done()

		// 显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	} ()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已完成
		// 保证每个goroutine一旦完成工作就调用Done
		defer wg.Done()

		// 显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	} ()

	// 等待 goroutine结束
	fmt.Println("Waiting To Finish!")
	wg.Wait()

	fmt.Println("Terminating Program")



}

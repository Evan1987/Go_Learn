package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// 显示逻辑处理器的切换情况，上一个示例太过简单导致还没切换就运行完了
func main() {

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 计数加2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	fmt.Println("Create Goroutines")

	// 实际运行时会发现调度器在不断切换线程上的程序
	go printPrime("A")
	go printPrime("B")


	// 等待goroutine结束
	fmt.Println("Wait to Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()

	next:  // 寻找5000以内质数
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner ++ {
			if outer % inner == 0 {
				continue next  // 表示continue外层循环，循环的标签
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
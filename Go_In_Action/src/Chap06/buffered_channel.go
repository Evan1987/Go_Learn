package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

const (
	numberGoroutines 	= 4   // 要使用的goroutine数量
	taskLoad 			= 10  // 要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func worker(tasks chan string, worker int) {
	// 通知函数已经返回
	defer wg.Done()

	for {
		task, ok := <- tasks
		if !ok {
			// 意味着通道已经空了，并且已经关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动goroutine来管理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr ++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post ++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 通道关闭，只是不能写入，还是可以继续读取缓冲值
	close(tasks)

	wg.Wait()
}

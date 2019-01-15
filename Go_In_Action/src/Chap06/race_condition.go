package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg2 sync.WaitGroup


// 竞争状态
func main() {

	runtime.GOMAXPROCS(1)

	wg2.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Println("Final Counter", counter)  // 逻辑处理器为1时，counter为2；其余为4
}

func incCounter(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count ++ {
		value := counter

		// 当前 goroutine从线程退出，并放回队列，强制调度
		runtime.Gosched()

		// 增加本地value变量的值
		value ++

		// 将该值保存回counter
		counter = value
	}
}
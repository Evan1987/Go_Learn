package main


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <- court
		if !ok {
			// 如果通道关闭了，我们就赢了
			fmt.Printf("Player %s Won \n", name)
			return
		}

		// 选择随机数，然后用这个随机数来判断我们是否会丢球
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数 + 1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}

func main() {
	court := make(chan int)  // 无缓冲通道

	wg.Add(2)

	// 启动两个通道
	go player("Nadal", court)
	go player("Djokovic", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}


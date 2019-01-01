package main

import (
	"./search"
	_ "./matchers"
	"log"
	"os"
)

// init 在 main之前调用
func init() {
	// 日志输出到标准输出
	log.SetOutput(os.Stdout)
}


func main(){
	// 使用特定项做搜索
	search.Run("president")
}


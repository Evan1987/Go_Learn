package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	x := duration(14)
	fmt.Println(x.pretty())
	//fmt.Println(duration(14).pretty())  // 错误！获取不到duration(14)的地址
}
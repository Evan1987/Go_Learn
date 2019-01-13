package main

import (
	"./counter"
	"./entities"
	"fmt"
)
func main() {
	counter := counter.New(10)
	fmt.Println(counter)

	//u := entities.User{"Bill", "bill@gmail.com"}  // error email未知
	//fmt.Printf("User: %v\n", u)

	a := entities.Admin{}  // 由于user未公开，则无法直接初始化
	a.Name = "Bill"
	a.Email = "bill@gmail.com"
	a.Rights = 10
	fmt.Printf("User: %v\n", a)
}

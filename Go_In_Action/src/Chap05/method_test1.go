package main

import "fmt"

type user struct {
	name 	string
	email 	string
}

// 使用 值接收者 实现一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}

// 使用 指针接收者 实现一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// go既允许使用值，又允许使用指针来调用方法

	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@live.cn"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@163.com")
	lisa.notify()  // (*lisa).notify()
}






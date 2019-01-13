package main

import "fmt"

type notifier2 interface {
	notify()
}

type userInfo struct {
	name string
	email string
}

func (u *userInfo) notify() {
	fmt.Printf("send user email to %s<%s>\n", u.name, u.email)
}

type adminInfo struct {
	userInfo
	level string
}

// adminInfo 外部类型也实现了notify方法
func (a *adminInfo) notify() {
	fmt.Printf("send admin email to %s<%s>\n", a.name, a.email)
}


func main() {
	ad := adminInfo{userInfo{"john smith", "john@gmail.com"}, "super"}
	send(&ad)  // 调用外部类型的方法
	send(&ad.userInfo)  // 调用内部类型的方法
}

func send(n notifier2) {
	n.notify()
}


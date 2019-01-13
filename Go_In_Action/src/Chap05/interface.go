package main

import "fmt"

// 任何实现了notify方法的对象都implement了notifier接口
type notifier interface {
	notify()
}

type user1 struct {
	name string
	email string
}

type user2 struct {
	id int
	sex string
}

// 利用指针接收者实现了notify方法
func (u *user1) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 利用值接收者实现了notify方法
func (u user2) notify() {
	fmt.Printf("No %d user's sex is %s\n", u.id, u.sex)
}


// 任何一个实现了notifier接口的值都可以传入
func sendNotification(n notifier) {
	n.notify()
}


func main() {
	u1 := user1{"Bill", "bill@gmail.com"}
	u2 := user2{1082, "female"}
	//sendNotification(u1)  // 错误！指针接收者只有指针实现了接口
	sendNotification(&u1)

	sendNotification(u2)  // 正确！值接收者，值与指针都实现了接口
	sendNotification(&u2)  // 正确！值接收者，值与指针都实现了接口
}

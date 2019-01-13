package main

import "fmt"

type notifier_ interface {
	notify()
}

type user_ struct {
	name string
	email string
}

func (u *user_) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name string
	email string
}

func (a *admin) notify() {
	fmt.Printf("sending user email to %s<%s>\n", a.name, a.email)
}


// 嵌入类型
type admin_ struct {
	user_
	level string
}

func main() {
	bill := user_{"bill", "bill@gmail.com"}
	lisa := admin{"lisa", "lisa@gmail.com"}
	ad := admin_{user_{"john smith", "john@gmail.com"}, "super"}

	sendNotification_(&bill)
	sendNotification_(&lisa)

	ad.notify()  // equal to ad.user_.notify()
	// 既然ad已经通过嵌入user_实现了notify方法，也就相当于实现了notifier接口，相当于sendNotification_(&ad.user_)
	sendNotification_(&ad)

}

// 多态函数
func sendNotification_(n notifier_) {
	n.notify()
}


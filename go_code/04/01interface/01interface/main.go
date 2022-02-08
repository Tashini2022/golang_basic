package main

import "fmt"

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵～")
}

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪～")
}

type sayer interface {
	//	接口：定义一个抽象类型，只要实现say（）就能称为sayer类型
	say()
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("滚一边儿去！")
}

//	不管是什么，只要执行da()这个函数就会调用say()
func da(arg sayer) {
	arg.say()
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{
		name: "它是你",
	}
	da(p1)
}

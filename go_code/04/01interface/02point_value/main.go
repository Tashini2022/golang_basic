package main

import "fmt"

//	接口可以嵌套其他接口
type animal interface {
	sayer
	mover
}

type sayer interface {
	say()
}
type mover interface {
	move()
}

//	使用指针接受者和值接受者的区别
type person struct {
	name string
	age  int
}

//
func (p person) say() { //值类型的接受者
	fmt.Printf("%s会说话!\n", p.name)
}

func (p *person) move() { //指针类型的接受者
	fmt.Printf("%s会运动!\n", p.name)
}

func main() {
	var s sayer   //	值类型
	var m mover   //	指针类型
	p1 := person{ //	p1是值类型的接受者
		name: "它是你",
		age:  18,
	}
	p2 := &person{ //	p2是指针类型的接受者
		name: "Tashini",
		age:  18,
	}

	s = p1
	s.say()
	fmt.Println(s)
	s = p2
	s.say()
	fmt.Println(s) //	s 是值类型，可以接受指针（P2）
	// m=p1			//	m 是指针类型，无法接收值（P1）
	m = p2
	m.move()
	fmt.Println(m)

	var a animal	//	a 是一个animal类型的接口
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}

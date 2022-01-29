package main

import "fmt"

//	嵌套结构体字段冲突

type Email struct {
	addr   string
	update string
}
type Address struct {
	city   string
	update string
}
type Person struct {
	name    string
	gender  string
	age     int8
	Address //	嵌套匿名结构体
	Email   //	嵌套另一个匿名结构体
}

func main() {
	p1 := Person{
		name:   "它是你",
		gender: "男",
		age:    27,
		Address: Address{
			update: "2022-01-10",
			city:   "贵阳",
		},
		Email: Email{
			addr:   "Tashini@126.com",
			update: "2020-12-12",
		},
	}
	fmt.Printf("p1=%#v ,type:%T\n", p1, p1)
	fmt.Println(p1.name, p1.gender, p1.age, p1.Address)
	fmt.Println(p1.Address.update) //	可通过嵌套的结构体字段访问其内部字段
	fmt.Println(p1.Email.update)
	// fmt.Println(p1.update)      //	update为两个嵌套结构体的同名字段，故不可直接访问
}

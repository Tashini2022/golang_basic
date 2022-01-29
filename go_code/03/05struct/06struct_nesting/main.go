package main

import "fmt"

//	结构体的嵌套

type Address struct {
	Province string
	city     string
}
type Person struct {
	name    string
	gender  string
	age     int8
	Address //	嵌套Address结构体，可匿名
}

func main() {
	p1 := Person{
		name:   "它是你",
		gender: "男",
		age:    27,
		Address: Address{
			Province: "贵州",
			city:     "贵阳",
		},
	}
	fmt.Printf("p1=%#v ,type:%T\n", p1, p1)
	fmt.Println(p1.name, p1.gender, p1.age, p1.Address)
	fmt.Println(p1.Address.city) //	通过嵌套的结构体字段访问其内部字段，
	fmt.Println(p1.city)         //	直接访问匿名结构体字段
}

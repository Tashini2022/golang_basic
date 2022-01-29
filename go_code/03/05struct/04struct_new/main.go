package main

import "fmt"

//	结构体构造函数：构造一个结构体实例化函数
type person struct {
	name, city string //同类型可以写在同一行
	age        int8
}

func newPerson(name, city string, age int8) *person { //结构体是值类型
	return &person{ //字段较多时会使用指针减少程序开销
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p1 := newPerson("它是你", "北京", 22)
	fmt.Printf("p1=%#v ,type:%T\n", p1, p1)
}

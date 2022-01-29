package main

import "fmt"

//	1。定义结构体
type person struct {
	name, city string //同类型可以写在同一行
	age        int8
}

func main() {
	//	结构体基本实例化
	var p person
	p.name = "它是你"
	p.city = "上海"
	p.age = 25
	fmt.Printf("p=%#v\n", p)
	fmt.Println(p.name, p.city, p.age)

	//	匿名结构体
	var usr struct {
		name string
		age  int8
	}
	usr.name = "它是你"
	usr.age = 100
	fmt.Printf("usr=%#v\n", usr)
	fmt.Println(usr.name, usr.age)

	//	结构体匿名字段
	type person1 struct {
		string //没有名称，所以同类型必须唯一
		int8
	}
	p1 := person1{
		"它是你",
		29,
	}
	fmt.Printf("p=%#v\n", p1)
	fmt.Println(p1.string, p1.int8)

}

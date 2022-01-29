package main

import "fmt"

type person struct {
	name, city string //同类型可以写在同一行
	age        int8
}

//结构体指针
func main() {
	var p1 = new(person)
	fmt.Printf("p1=%#v ,type:%T\n", p1, p1)
	//两种实例化方式
	/*(*p1).name = "它是你"
	(*p1).city = "上海"
	(*p1).age = 25*/
	p1.name = "它是你"
	p1.city = "上海"
	p1.age = 25
	fmt.Printf("p=%#v\n", p1)
	fmt.Println(p1.name, p1.city, p1.age)

	//取结构体地址进行实例化
	p2 := person{}
	fmt.Printf("p2=%#v ,type:%T\n", p2, p2)
	p3 := &person{}
	fmt.Printf("p3=%#v ,type:%T\n", p3, p3)
	p3.name = "它是你a"
	p3.city = "上海a"
	p3.age = 20
	fmt.Printf("p3=%#v ,type:%T\n", p3, p3)
}

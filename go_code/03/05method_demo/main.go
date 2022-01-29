package main

import "fmt"

//方法的定义实例

//Person 是一个结构体
type Person struct {
	name string
	age  int8
}

//	NewPerson 是一个Peraon类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age}
}

//Dream 是为Person类型的定义方法
func (p Person) Dream() {
	fmt.Printf("%v想看: 落霞与孤鹜齐飞，秋水共长天一色。\n", p.name)
}

//SetAge 是一个修改年龄的方法
//指针接受者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

//给任意类型的变量定义方法
//无法为系统类型int定义方法，但可以基于int类型定义myInt，然后为myInt定义方法
type myInt int

func (m myInt) sayhi() {
	fmt.Println("hi")
}

//SetAge2 是一个利用值类型接受者修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("它是你", int8(18))
	//(*p1).Dream 亦可
	p1.Dream()
	fmt.Println(p1.age)
	p1.SetAge(int8(23))   //	调用指针类型接受者方法修改
	fmt.Println(p1.age)   //	23
	p1.SetAge2(int8(100)) //	调用值类型接受者方法修改
	fmt.Println(p1.age)   //	23？

	var m1 myInt
	m1.sayhi()
}

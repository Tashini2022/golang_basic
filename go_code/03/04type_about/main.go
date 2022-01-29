package main

import "fmt"

//	1.自定义类型
type Myint int

//	2.类型别名
// 	Newint int类型别名
type Newint = int

func main() {
	var i Myint
	fmt.Printf("type:%T,value:%v\n", i, i)
	var y Newint
	fmt.Printf("type:%T,value:%v\n", y, y)

}

package main

import (
	"fmt"
)

// //匿名函数与闭包
// //定义一个函数，返回值是一个函数
// func a(name string) func() {
// 	return func() {
// 		fmt.Println("你好", name)
// 	}
// }

// func main() {
// 	r := a("它是你")
// 	r()//	“你好，它是你”
// }
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(123)
	fmt.Println(f1(1), f2(5))  //124 119
	fmt.Println(f1(32), f2(7)) //151 144
	fmt.Println(f1(21), f2(9)) //165 156
}

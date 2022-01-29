package main

import "fmt"

//申明函数，可以使用不同的参数

//申明可变参数函数，可变参数实际为切片
func calc(x int, y ...int) int {
	ret := x
	for _, arg := range y {
		ret = ret + arg
	}
	return ret
}

//函数作为参数
func add(x, y int) int {
	return x + y
}
func cal(x, y int, op func(x, y int) int) int { //参数类型简写
	return op(x, y)
}

func main() {

	//函数作为变量和参数
	aa := calc

	fmt.Println(cal(200, 100, add))
	fmt.Println(aa(200, 100))
}

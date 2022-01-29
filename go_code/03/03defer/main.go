package main

import "fmt"

//defer:延迟执行
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func b() {
	defer func() {
		err := recover() //只申明变量，未申请内存地址
		if err != nil {
			fmt.Println("func b error!")
		}
		panic("panic in b")
	}() //（）表示立即执行
}
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) //1先执行内部A,AA暂存。通常和recover()同时使用
	x = 10
	defer calc("BB", x, calc("B", x, y)) //2执行B，3执行BB，4执行AA
	y = 20                               //为使用d，efer使用时必须先申明变量
}

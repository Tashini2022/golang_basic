package main

import "fmt"

/*
func main() {
	//指针
	//1  &：取得内存地址
	var n = 23
	var p = &n
	fmt.Println(p)
	fmt.Printf("p = %x(%T)\n", p, p) //*int:int类型的指针
	//2  *：根据地址取值
	fmt.Printf("*p =%d(%T)\n", *p, *p)
	//new()函数操作
	var a1 *int       //定义一个指针a1
	fmt.Println(a1)   //a1默认值为nil所以*a1无法执行
	var a2 = new(int) //new()函数申请一个内存地址
	fmt.Println(a2, *a2)
	*a2 = 1212
	fmt.Println(a2, *a2)
	//
}*/
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}

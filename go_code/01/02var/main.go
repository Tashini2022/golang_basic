package main

import "fmt"

func main() {
	var ( //var 为变量，const 为常量
		n1   int
		isOk bool //中间单词首字母大写
	)
	/*批量申明变量
	全局变量可不使用
	*/
	var s = "落霞与孤鹜齐飞，秋水共长天一色"
	//申明同时赋值（局部变量必须使用）编译器会自动识别变量类型
	fmt.Println(s)
	n1 = 2022
	fmt.Printf("今年是%d年!", n1)
	isOk = false
	fmt.Print("今年是2021年？\n", isOk)
	ss := "2022年1月17日" //简短申明，只能用于局部变量
	//匿名变量“_”多用于占位（会被编译器丢弃）
	fmt.Println(ss)
}

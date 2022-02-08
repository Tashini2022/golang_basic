package main

import "fmt"

//常量申明
const pi = 3.1415926

//批量申明
const (
	n1 = 100
	n2 = 250
	s1 = "300"
	s2
	s3
	//批量申明时如果某一行内有赋值，则默认与上一行相同
)

//常量计数器: iota
//在const关键字出现时将重置为0，每新增一行常量申明，则iota 增加1；

const (
	b1 = iota
	b2 = iota
	_  = iota
	b3 = iota
)

//插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

//多个常量、声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //"<<"左移符号
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("pi =", pi, "n1 =", n1, "n2 =", n2, "s1 =", s1, "s2 =", s2, "s3 =", s3)
	fmt.Println("b1 =", b1, "b2=", b2, "b3=", b3)
	fmt.Println("c1 =", c1, "c2=", c2, "c3=", c3, "c4 =", c4)
	fmt.Println("d1 =", d1, "d2=", d2, "d3=", d3, "d4 =", d4)
}

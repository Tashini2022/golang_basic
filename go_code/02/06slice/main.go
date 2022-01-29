package main

import "fmt"

//切片
func main() {
	//切片定义
	var s1 []int
	var s2 []string
	fmt.Println("s1:", s1, "s2:", s2)
	fmt.Println(s1 == nil, s2 == nil)

	//切片初始化
	s1 = []int{1, 2, 3, 4, 5}
	s2 = []string{"北京", "南京", "西安", "上海"}
	fmt.Println("s1:", s1, "s2:", s2)
	fmt.Println(s1 == nil, s2 == nil)

	//长度和容量
	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))

	//由数组定义切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	as1 := a1[2:6] //[3 4 5 6]
	as2 := a1[:5]  //[0:5]
	as3 := a1[3:]  //[3:len(a1)]
	as4 := a1[:]   //[0:len(a1)]
	fmt.Println(as1, as2, as3, as4)

	//切片容量是指从切片的第一个元素到数组最后一个元素
	//切片的长度是他本身的长度
	fmt.Println("len(as1):", len(as1), "cap(as1):", cap(as1))
	fmt.Println("len(as2):", len(as2), "cap(as2):", cap(as2))
	fmt.Println("len(as3):", len(as3), "cap(as3):", cap(as3))

	//切片再切割，切片为引用类型，更改数组切片会更改
	as22 := as2[3:]
	fmt.Println("len(as22):", len(as22), "cap(as22):", cap(as22), as22)
	a1[3] = 123
	fmt.Println(as1, as2, as3, as4)
	fmt.Println("len(as22):", len(as22), "cap(as22):", cap(as22), as22)
}

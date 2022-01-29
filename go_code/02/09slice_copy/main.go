package main

import "fmt"

func main() {
	x1 := [...]int{1, 2, 3, 4, 5} //数组
	s1 := x1[:4]                  //由数组得到切片
	s2 := s1                      //为s2赋值
	fmt.Printf("s1=%v,cap(s1)%d: ,len(s1)%d: \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,cap(s2)%d: ,len(s2)%d: \n", s2, len(s2), cap(s2))
	var s3 = make([]int, 4)
	copy(s3, s1) //copy切片
	s1[2] = 200
	fmt.Printf("s1=%v,cap(s1)%d: ,len(s1)%d: \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,cap(s2)%d: ,len(s2)%d: \n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v,cap(s3)%d: ,len(s3)%d: \n", s3, len(s3), cap(s3))

	//将某一个元素从切片中删除
	s1 = append(s1[:2], s1[3:]...) //只能通过追加的方式删除切片元素
	fmt.Printf("s1=%v,cap(s1)%d: ,len(s1)%d: \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,cap(s2)%d: ,len(s2)%d: \n", s2, len(s2), cap(s2))//底层数组的改变导致s2改变[1 2 4 4]
	fmt.Printf("s3=%v,cap(s3)%d: ,len(s3)%d: \n", s3, len(s3), cap(s3))
	fmt.Printf("x1=%v,\n", x1) //切片的删除修改底层数组[1 2 4 4 5]

}

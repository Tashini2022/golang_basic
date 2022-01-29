package main

import "fmt"

func main() {
	//make()函数创建切片
	s1 := make([]int, 5, 10)
	fmt.Println(s1, "len(s1):", len(s1), "cap(s1):", cap(s1))
	s2 := make([]int, 5)
	fmt.Println(s2, "len(s2):", len(s2), "cap(s2):", cap(s2))

	//切片赋值
	var s3 = []int{1, 2, 3, 4}
	var s4 = s3 //指向同一个底层数组
	fmt.Println(s3, s4)
	s3[1] = 222
	fmt.Println(s3, s4)

	//切片遍历
	for i := 0; i < len(s3); i++ { //方案一
		fmt.Println(i, s3[i])
	}
	for k, v := range s3 { //方案二
		fmt.Println(k, v)
	}
}

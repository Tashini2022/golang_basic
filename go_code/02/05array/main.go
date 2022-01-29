package main

import "fmt"

func main() {
	var s1 [3]bool //数组长度[]与类型是数组故有属性
	var s2 [5]bool //不同数组不可比较
	var s3 [4]int
	fmt.Printf("s1: %T\t,s2: %T\t, s3: %T\n", s1, s2, s3)

	//数组初始化（默认值bool:false , int和float:0，string:""空字符）
	a1 := [3]bool{true, false, true}  //方式一
	a2 := [...]int{1, 3, 5, 6, 2, 4}  //方式二
	a3 := [6]string{0: "北京", 5: "上海"} //方式3
	fmt.Print("a1=", a1, "\ta2=", a2, "\ta3=", a3)

	//数组遍历
	for i := 0; i < len(a2); i++ { //for循环遍历数组
		fmt.Println(a2[i])
	}

	for k, v := range a3 { //for range循环遍历数组
		fmt.Println(k, v)
	}

	//二维数组
	var a11 [4][3]int
	a11 = [4][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
		[3]int{10, 11, 12},
	}
	for i := 0; i < len(a11); i++ { //for 遍历数组
		fmt.Println(a11[i])
		for j := 0; j < len(a11[i]); j++ {
			fmt.Println(a11[i][j])
		}
	}
	for _, v1 := range a11 { //for range 遍历数组
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//求出数组中和为特定值(7)的两个元素的下标
	for i := 0; i < len(a2)-1; i++ {
		for j := i + 1; j < len(a2); j++ {
			if a2[i]+a2[j] == 7 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

package main

import "fmt"

//if循环

func main() {
	//条件判断
	var s = 87
	if s >= 90 {
		fmt.Println("成绩优秀")
	} else if s >= 80 {
		fmt.Println("成绩良好")
	} else if s >= 70 {
		fmt.Println("成绩中等")
	} else if s >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	//特殊写法
	if s1 := 77; s1 >= 90 {
		fmt.Println("成绩优秀")
	} else if s1 >= 70 {
		fmt.Println("成绩良好")
	} else if s1 >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}
}

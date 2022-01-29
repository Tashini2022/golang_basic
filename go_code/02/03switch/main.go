package main

import "fmt"

func main() {
	var n = 3
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效数字")
	}

	//"switch"语句用于简化"else if"语句
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效数字")
	}

	//switch变式1
	switch k := 9; k {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("无效数字")
	}

	//switch变式2
	switch {
	case n < 0:
		fmt.Println("负数")
	case n > 0:
		fmt.Println("正数")
	}
}

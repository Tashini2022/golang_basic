package main

import "fmt"

func main() {
	var flag = false //利用j跳出多层for循环
	for k := 1; k <= 9; k++ {
		for j := 1; j <= k; j++ {
			if j == 4 {
				flag = true
				continue
			}
			fmt.Printf("%d x %d = %d\t", k, j, k*j)
		}
		if flag {
			flag = false
			fmt.Print("hahaha")
		}
		fmt.Printf("\n")
	}

	//goto语句 + label标签实现跳转
	for k := 1; k <= 9; k++ {
		for j := 1; j <= k; j++ {
			if j == 4 {
				goto XX
			}
			fmt.Printf("%d x %d = %d\t", k, j, k*j)
		}

		fmt.Printf("\n")
	}

	//标签
XX:
	fmt.Print("不想写了")
}

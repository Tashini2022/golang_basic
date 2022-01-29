package main

import "fmt"

//for 循环
func main() {
	var n = 10
	var sum = 0
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		sum = sum + i
	}
	fmt.Printf("sum=%d\n", sum)

	//变种1,省略初始语句
	var i = 5
	var s1 = 0
	for ; i <= n; i++ {
		fmt.Println(i)
		s1 = s1 + i
	}
	fmt.Printf("sum=%d\n", s1)

	//变种2,省略结束语句
	var s2 = 0
	var y = 0
	for y <= n {
		fmt.Println(y)

		s2 = s2 + i
		y++
	}
	fmt.Printf("sum=%d\n", s2)

	//乘法表
	for k := 1; k <= 9; k++ {
		for j := 1; j <= k; j++ {
			fmt.Printf("%d x %d = %d\t", k, j, k*j)
		}
		fmt.Printf("\n")
	}

	//当k==5时跳出for循环，不再执行
	for k := 9; k >= 1; k-- {
		if k == 5 {
			fmt.Printf("\n")
			break
		}
		for j := 1; j <= k; j++ {

			fmt.Printf("%d x %d = %d\t", k, j, k*j)
		}
		fmt.Printf("\n")
	}

	//当k==5时跳过本次for循环，继续下一次循环
	for k := 9; k >= 1; k-- {
		if k == 5 {
			fmt.Printf("\n")
			continue
		}
		for j := 1; j <= k; j++ {

			fmt.Printf("%d x %d = %d\t", k, j, k*j)
		}
		fmt.Printf("\n")
	}
}

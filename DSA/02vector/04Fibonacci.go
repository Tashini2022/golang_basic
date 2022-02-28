package main

// 递归
// func Fibonacci1(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return Fibonacci1(n-2) + Fibonacci1(n-1)
// }

// 迭代
func Fibonacci2(n int) int {
	if n < 2 {
		return n
	} else {
		g, f := 1, 0 // 计算斐波纳契数列的第n项
		for i := 2; i < n; i++ {
			g = g + f
			f = g - f
		}
		return g
	}
}

package main

// 计算n个数的和
func Sum(n int) int {
	if n < 2 {
		return n
	} else {
		return Sum(n-1) + n /* 函数调用自身 */
	}
}

// 计算n的阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

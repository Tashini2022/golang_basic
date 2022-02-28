package main

// 对数组的第lo—hi和元素求和
func Sum1(n []int, lo, hi int) int {
	if lo == hi {
		return n[hi]
	} else {
		mi := (lo + hi) / 2
		return Sum1(n, lo, mi) + Sum1(n, mi+1, hi) /* 函数调用自身 */
	}
}

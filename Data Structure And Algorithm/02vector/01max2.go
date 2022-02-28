package main

import "fmt"

// 半开半闭区间
// 在给定的区间[lo,hi)中找出最大的两个元素，
// A[x1]>=A[x2]
var A = []int{2, 5, 8, 6, 9, 4, 7, 10, 1, 11, 3, 4, 5, 2} //9,7

// 方法一
func Max1(lo, hi int) (x1, x2 int) {
	x1 = lo
	for i := lo + 1; i < hi; i++ {
		if A[x1] < A[i] {
			x1 = i
		}
	}
	x2 = lo
	for i := lo + 1; i < x1; i++ {
		if A[x2] < A[i] {
			x2 = i
		}
	}
	for i := x1 + 1; i < hi; i++ {
		if A[x2] < A[i] {
			x2 = i
		}
	}
	return
}

// 方法二
func Max2(lo, hi int) (x1, x2 int) {
	x1, x2 = lo, lo+1 // 交换使得A[x1]>A[x2]
	if A[x1] < A[x2] {
		x1, x2 = x2, x1
	}
	for i := lo + 2; i < hi; i++ {
		if A[i] > A[x2] {
			x2 = i //将数值小的替换
			if A[x1] < A[x2] {
				x1, x2 = x2, x1
			}
		}
	}
	return
}

// A[x1]>=A[x2]
// 递归+分治
func Max3(A []int, lo, hi int) (x1, x2 int) {
	if lo+2 == hi {
		x1, x2 = lo, lo+1  // 切片中只有两个元素
		if A[x1] < A[x2] { // 如果A[x1] < A[x2]则交换下标保证A[x2] < A[x1]
			x1, x2 = x2, x1
		} // else A[x2] < A[x1]
	} else if lo+3 == hi { // 切片中只有三个元素
		x1, x2 = lo, lo+1  // 先赋值
		if A[x2] < A[x1] { // 在A[x2] < A[x1]时比较A[x2]和另外一个数
			if A[x2] < A[hi-1] { // A[x2] < A[hi-1]时替换x2的值
				x2 = hi - 1
				if A[x1] < A[x2] { // 然后再比较替换后的数字
					x1, x2 = x2, x1
				}
			}
		} else { // A[x1] < A[x2]时
			x1, x2 = x2, x1      // 交换下标，保证A[x2] < A[x1]
			if A[x2] < A[hi-1] { // 再用小元素和另一个元素比较
				x2 = hi - 1
				if A[x1] < A[x2] {
					x1, x2 = x2, x1
				}
			}
		}
	} else { // 切片内部元素大于三个
		mi := int((lo + hi) >> 1)
		x1l, x2l := Max3(A, lo, mi)
		x1r, x2r := Max3(A, mi, hi)
		if A[x1l] > A[x1r] {
			if A[x2l] > A[x1r] { //x1l>x2l>x1r
				x1, x2 = x1l, x2l
			} else { //x1l>x1r>x2l
				x1, x2 = x1l, x1r
			}
		} else { //x1r>x1l
			if A[x2r] > A[x1l] { //x1r>x2r>x1l
				x1, x2 = x1r, x2r
			} else { //x1r>x1l>x2r
				x1, x2 = x1r, x1l
			}
		}
	}
	return
}
func main() {
	fmt.Println(Max1(3, 13))
	fmt.Println(Max2(3, 13))
	fmt.Println(Max3(A, 5, 10))
}

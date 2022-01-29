package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	for i := 10; i >= 0; i-- {
		a = append(a, i)
	}
	fmt.Println(a)
	sort.Ints(a[:])//排序
	fmt.Println(a)
}

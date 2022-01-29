package main

import "fmt"

func main() {
	//append追加切片,必须用一个变量接受追加内容
	var s1 = []string{"北京", "上海", "广州"}
	fmt.Printf("s1=%v，len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "深圳")
	fmt.Printf("s1=%v，len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
	var s11 = []string{"南京", "杭州", "成都"}
	s1 = append(s1, s11...) //...表示拆开s11,否则无法追加
	fmt.Printf("s1=%v，len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "武汉", "西安")
	fmt.Printf("s1=%v，len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
}

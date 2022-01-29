package main

import (
	"fmt"
	"strings"
	"unicode"
)

//双引号表示字符串，单引号表示字符

func main() {
	//字符串
	//转义字符
	//  \"表示双引号，\'表示单引号，\\表示反斜杠，
	//  \r表示返回本行首，\n表示换行，\t表示制表
	s1 := "\"真美啊！\""
	fmt.Println(s1)

	//多行字符串（反引号esc键下方）
	s2 := `
	天苍苍
	野茫茫
	风吹草低见牛羊
	`
	fmt.Println(s2)

	//字符串拼接
	ss := s2 + s1
	fmt.Println(ss)
	fmt.Printf("%s%s\n", s2, s1)
	ss1 := fmt.Sprintf("%s%s\n", s2, s1)
	fmt.Println(ss1)

	//分割
	s3 := "/home/tashini/Golang/goproject/src/go_code/01/06String/main.go"
	ss3 := strings.Split(s3, "/") // 分割s3
	fmt.Println(ss3)              // [home tashini Golang goproject src go_code 01 06String main.go]

	//拼接join操作
	fmt.Println(strings.Join(ss3, "++"))

	//包含
	fmt.Println(strings.Contains(s3, "ta")) //判断是否包含字符串，返回true或者false
	fmt.Println(strings.Contains(s3, "ts"))

	//前缀/后缀判断，返回true或者false
	fmt.Println(strings.HasPrefix(s3, "/h"))
	fmt.Println(strings.HasSuffix(s3, "main.go"))

	//计算子字符串出现的位置（第一次或者最后一次）
	var s4 = "21abcdcba12"
	fmt.Println(strings.Index(s4, "ab"))
	fmt.Println(strings.LastIndex(s4, "1"))

	var count int
	var str = "hello 世界"
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}

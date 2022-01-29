package main

import "fmt"

func main() {
	var i1 = 109                    //十进制数
	fmt.Printf("十进制数 \t%d\n", i1)   //输出为十进制数
	fmt.Printf("输出为二进制数\t%b\n", i1) //输出为二进制数
	fmt.Printf("输出为八进制数\t%o\n", i1) //输出为八进制数
	fmt.Printf("输出为十六进制数%x\n", i1)  //输出为十六进制数
	fmt.Printf("输出为字符%c\n", i1)     //输出为字符

	// 八进制数，0 开头
	i2 := 077
	fmt.Printf("i2 = %d\n", i2)

	//十六进制数，0x开头
	i3 := 0xff
	fmt.Printf("i3 = %d\n", i3)

	//查看变量类型 "%T"
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9) //明确为int8类型，否则为默认int类型
	fmt.Printf("%T\n", i4)

	var s = "hello 世界!"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)  //查看内容
	fmt.Printf("%#v\n", s) //查看内容
}

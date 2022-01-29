package main

import "fmt"

func main()  {

	//len求字节数
	var s1 = "abcdef!"
	s2 := "秋水共长天一色"
	s3 := "abc一二三"
	n1 := len(s1)
	n2 := len(s2)
	fmt.Println("n1 = ", n1 , "n2 = ", n2)

	//强制转型之后才能处理汉字
	ss3 := []rune(s3)
	for i := 0 ; i < len([]rune(s3)); i++ {
		fmt.Println(ss3[i])
		fmt.Printf("%c\n",ss3[i])
	}
	//修改字符串，先转化为rune（字节），再修改
	ss3[2] , ss3[5]= 'a' ,'一'
	fmt.Println(string(ss3))//把rune切片强制转型为string类型
	fmt.Printf("s3[0]:%T\n" , s3[0])
	fmt.Printf("ss3[0]:%T\n" , ss3[0])
}
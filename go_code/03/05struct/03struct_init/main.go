package main

import "fmt"

type person struct {
	name, city string //同类型可以写在同一行
	age        int8
}

func main() {
	//	1.键值对初始化
	p3 := person{
		name: "它是你", //可以只初始化一部分字段
		city: "武汉",
		age:  18,
	}
	fmt.Printf("p3=%#v ,type:%T\n", p3, p3)
	//	2.值的列表进行初始化
	p4 := person{ //须按照字段顺序全部写出
		"它是你",
		"成都",
		20,
	}
	fmt.Printf("p4=%#v ,type:%T\n", p4, p4)
}

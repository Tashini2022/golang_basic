package main

import "fmt"

func main() {
	//申明map()的同时进行初始化
	var m2 = map[int]string{
		1: "1223",
		2: "2334",
		3: "4556",
	}
	//判断某个键值是否存在map()中
	value, ok := m2[3]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("存在此数据", value, ok)
	} else {
		fmt.Println("不存在此数据")
	}

	var m1 map[string]int
	fmt.Println(m1 == nil, len(m1)) //默认为nil（未申请内存）
	m1 = make(map[string]int)       //map()初始化方式一，可在运行过程中动态加载内存（但为了程序运行效率，不建议这么做）
	fmt.Println(m1, len(m1), m1 == nil)
	m1["世界"] = 2022
	m1["孔子"] = 2000
	m1["李白"] = 1212
	fmt.Println(m1, len(m1))

	//删除map，如果对象不存在，则不执行亦不报错
	delete(m1, "李白")
	delete(m1, "2000")
	fmt.Println(m1)

	//map()遍历，k,v都可单独遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

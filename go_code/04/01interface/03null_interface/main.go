package main

import "fmt"

//	空接口
//	没有定义任何方法的接口，空接口可以储存任意值
//	空接口一般不需要提前定义

func main() {
	var a interface{}
	a = 18
	fmt.Println(a)
	a = "Tashini"
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 空接口的应用
	// 1 作为函数的参数
	// 2 作为map的value
	var m = make(map[string]interface{}, 16)
	m["name"] = "Tashini"                       //	字符串
	m["age"] = 27                               //	int
	m["hobby"] = []string{"drawing", "dancing"} //	切片
	fmt.Println(m)

	// 类型断言
	// 猜测类型不正确时，ok值为false(bool)，ret为类型零值
	var y interface{}
	y = "Tashini"
	ret, ok := y.(int)
	if !ok {
		fmt.Println("y不是int类型！")
	} else {
		fmt.Println("猜测正确，ret=", ret)
	}

	// 使用switch进行类型断言
	switch t := y.(type) {
	case string:
		fmt.Printf("y是string类型，值为:%v\n", t)
	case int:
		fmt.Printf("y是int类型，值为:%v\n", t)
	case bool:
		fmt.Printf("y是bool类型，值为:%v\n", t)
	default:
		fmt.Printf("未知类型，y的值为:%v\n", t)
	}
}

package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 为了遍历结构体的所有方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.NumMethod()) // 打印结构体的方法数量

	v := reflect.ValueOf(x)
	// 由于需要具体的方法去调用，故使用的是值v
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 调用方法
	}

	// 通过方法名获取结构体的方法
	// t.MethodByName("Sleep") 返回值为（Method,bool）
	// v.MethodByName("Sleep") 返回值是 value
	fmt.Println(v.MethodByName("Sleep").Type())

}
func main() {
	stu1 := student{
		Name:  "它是你",
		Score: 90,
	}

	// 通过反射去获取结构体中的所有信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v\ttype:%v\n", t.Name(), t.Kind())
	// for循环遍历结构体所有字段
	for i := 0; i < t.NumField(); i++ {
		fieldobj := t.Field(i)
		fmt.Printf("name:%v\ttype:%v\ttag:%v\n", fieldobj.Name, fieldobj.Type, fieldobj.Tag)
		// 显示tag
		fmt.Println(fieldobj.Tag.Get("json"), "\t", fieldobj.Tag.Get("ini"))
	}

	// 根据名字去取结构体内的字段
	fieldobj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v\ttype:%v\ttag:%v\n", fieldobj2.Name, fieldobj2.Type, fieldobj2.Tag)
	} else {
		fmt.Println("未找到Score字段")
	}

	// 调用方法
	printMethod(stu1)
}

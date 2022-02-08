package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 由于不知道别人调用该函数时会传入什么类型的变量
	// 方式一，通过类型断言
	// 方式二，借助反射
	obj := reflect.TypeOf(x)
	fmt.Printf("obj: %v\t,name: %v\tkind: %v\n", obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) // *reflect.rtype
}



type cat struct{}
type dog struct{}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	var c = cat{}
	var d = dog{}
	reflectType(c)
	reflectType(d)
	var e []int
	var f []string
	//切片，数组，Map等返回的.Name为空
	reflectType(e)
	reflectType(f)

}

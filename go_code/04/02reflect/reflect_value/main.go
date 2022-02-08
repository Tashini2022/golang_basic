package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value:%v\ttype:%T\n", v, v)
	// 如何得到一个传入时类型的变量
	k := v.Kind() // 拿到对应类型的种类
	switch k {
	case reflect.Float32: // 把反射取得的值转换成int32类型的变量
		ret := float32(v.Float())
		fmt.Printf("value:%v\ttype:%T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("value:%v\ttype:%T\n", ret, ret)
	}
}

// 利用反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// ELem()根据指针取值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.14159)
	}
}

func main() {
	var aa int32 = 100
	reflectValue(aa)
	var bb float32 = 1.123
	reflectValue(bb)
	var aaa int32 = 10
	reflectSetValue(&aaa) // 传参为拷贝所以应该用指针
	fmt.Println(aaa)
}

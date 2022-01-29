package main

import (
	"fmt"
	"strings"
)

func main() {
	//	元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) //	对切片进行初始化，但元素map未初始化
	fmt.Println(mapSlice)
	fmt.Println(mapSlice[0] == nil)
	mapSlice[0] = make(map[string]int, 8) //	对切片的第一个元素初始化
	mapSlice[0]["北京"] = 010               //	对切片的第一个元素进行赋值
	fmt.Println(mapSlice)

	//值为切片的map
	var sliceMap = make(map[string][]int, 8) //	对map进行初始化
	fmt.Println(sliceMap, len(sliceMap))
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v, sliceMap["中国"])
	} else {
		//sliceMap中没有中国这个键
		sliceMap["中国"] = make([]int, 8) //	完成了对切片的初始化
		sliceMap["中国"][0] = 123
		sliceMap["中国"][1] = 234
		sliceMap["中国"][2] = 345
		sliceMap["中国"][3] = 456
		sliceMap["中国"][4] = 567
	}
	//	for遍历切片
	for v, k := range sliceMap {
		fmt.Println(v, k)
	}
	fmt.Println(sliceMap)

	//	统计“How do you do”中每个单词出现的次数
	//	1.创建map计数
	var s = "How do you do"
	worldCount := make(map[string]int, 8)
	//	2.分割单词
	worlds := strings.Split(s, " ")
	//	3.遍历单词作统计
	for _, world := range worlds {
		_, ok := worldCount[world]
		if ok {
			worldCount[world]++
		} else {
			worldCount[world] = 1
		}
	}
	fmt.Println(worldCount)
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//	创建map(),并初始化
	var Scoremap = make(map[string]int, 50)
	//	添加50个键值对
	for i := 1; i <= 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		score := rand.Intn(100)
		Scoremap[key] = score
	}
	//	打印Scoremap
	for k, v := range Scoremap {
		fmt.Println(k, v)
	}

	//	按照key大小排序Scoremap
	//	1.取出key，并存入slice
	keys := make([]string, 0, 50) //	先创建切片
	for k := range Scoremap {     //	存入数据
		keys = append(keys, k)
	}
	//	2.对储存key的slice进行排序
	sort.Strings(keys)
	//	3.按照排序完成的slice打印
	for _, v := range keys {
		fmt.Println(v, Scoremap[v])
	}
}

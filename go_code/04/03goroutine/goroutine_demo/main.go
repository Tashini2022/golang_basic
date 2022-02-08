package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // sync.WaitGroup
func hello(x interface{}) {
	fmt.Printf("hello %v！\n", x)
	wg.Done() //通知wg计数器-1
}

func main() { // 开启一个goroutine去执行main函数

	// wg.Add(1)  // 计数器+1
	for i := 0; i < 1000; i++ { //开启10000个goroutine
		wg.Add(1)
		go hello(i) // 在函数名前加go关键字，开启一个goroutine去执行hello这个函数
	}

	fmt.Println("hello main!")
	// time.Sleep(time.Second)	// 等待goroutine执行完毕！！！不建议使用
	wg.Wait() // 等待goroutine执行完毕
}

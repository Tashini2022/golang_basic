package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // sync.WaitGroup

func main() { // 开启一个goroutine去执行main函数
	runtime.GOMAXPROCS(2)	// GOMAXPROCS 确定调用几个逻辑核心执行：默认为全部
	// wg.Add(1)  // 计数器+1
	for i := 0; i < 1000; i++ { //用goroutine开启匿名函数
		wg.Add(1)
		go func(x int) {
			fmt.Println("hello ", x)
			wg.Done()
		}(i) //如果不传参，闭包会导致死锁
	}

	fmt.Println("hello main!")
	// time.Sleep(time.Second)	// 等待goroutine执行完毕！！！不建议使用
	wg.Wait() // 等待goroutine执行完毕
}

package main

import "fmt"

func main() {
	var ch1 chan int // 需要使用make（）函数初始化之后才能使用
	ch1 = make(chan int, 1)	// 带缓冲区通道，异步通道
	// ch1:=make(chan int)  // 无缓冲区通道，同步通道
	//必须要有数据接收方才可往里面发送数据，否则会阻塞导致死锁
	ch1 <- 20		// 往通道中传入数据
	x := <-ch1		// 从通道中获取数据
	fmt.Println(x)
	// l:=len(ch1) //获取通道中元素数量
	// c:=cap(ch1) // 获取通道容量
	close(ch1)	// 非必须操作，垃圾回收机制会自动回收通道
}

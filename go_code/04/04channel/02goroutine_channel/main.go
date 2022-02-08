package main

import "fmt"

/*
两个goroutine
1. 生成0—100的数字发送到ch1中
2. 从ch1中获取数字，计算后发送到ch2中
*/

func f1(ch chan<- int) {	//chan<- 只写单向通道。用于函数参数，限制函数对通达的操作
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func f2(ch1 <-chan int, ch2 chan<- int) {	//chan<- 只写单向通道。<-chan 只读单向通道
	//从通道中取值的方式一
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		} else {
			ch2 <- tmp * tmp
		}
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)
	//从通道中取值的方式2
	for ret := range ch2 { 
		fmt.Println(ret)
	}

}

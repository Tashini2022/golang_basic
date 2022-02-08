package main

import (
	"fmt"
	"sync"
)

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁
	rwlock sync.RWMutex // 读写互斥锁
)

func add() {
	for i := 0; i < 500000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //释放锁
	}
	wg.Done()
}


func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

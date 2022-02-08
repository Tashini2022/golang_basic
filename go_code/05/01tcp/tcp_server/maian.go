package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(conn net.Conn){
	defer conn.Close()	// 处理完之后关闭连接
	for{
		reader:=bufio.NewReader(conn)
		var buf [1024]byte
		n,err:=reader.Read(buf[:])
		if err!=nil{
			fmt.Printf("read from conn failed,err:%v\n", err)
			break
		}
		recv:=string(buf[:n])
		fmt.Printf("接收到的数据是:%v\n", recv)
		conn.Write([]byte("ok!"))	// 给客户端回复
	}
}
func main() {
	// 1.启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	// 2.等待客户端建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}
	// 3.启动单独的goroutine处理连接
		go process(conn)
	}
	
}

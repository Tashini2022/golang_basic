package main

import (
	"fmt"
	"net"
)

// udp server demo

func main() {
	// 1.启动服务
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10000,
	})
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	// 2.等待客户端建立连接
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read from udp failed,err:%v\n", err)
			return
		}
		fmt.Printf("接收到的数据是:%v\n", string(data[:n]))
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("write to udp failed,err:%v\n", err)
			return
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client demo

func main() {
	// 1.与服务端建立连接
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10000,
	})
	if err != nil {
		fmt.Printf("dial failed,err:%v\n", err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	// 2.数据的发送与接收
	for {
		s, _ := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
			return
		}
		// 从服务端接受回复
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read from udp failed,err:%v\n", err)
			return
		}
		fmt.Printf("read from %v,mag:%v\n", addr, string(buf[:n]))
	}

}

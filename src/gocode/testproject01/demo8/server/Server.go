package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("创建一个服务端")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("2222")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("6666768")
		}

		go process(conn)
	}

}

func process(conn net.Conn) {

	byteCli := make([]byte, 1024)
	num, err := conn.Read(byteCli)
	if err != nil {
		fmt.Println(565656)
	}

	fmt.Printf("从客户端接收到的数据%v", string(byteCli[0:num]))

}

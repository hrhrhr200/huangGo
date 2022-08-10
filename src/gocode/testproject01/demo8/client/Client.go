package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("开启一个客户端！！！")

	//定义一个客户端进行消息消息接收
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接客户端异常")
		return
	}

	fmt.Printf("连接到的客服端为:%v\n", conn)

	content, err2 := bufio.NewReader(os.Stdin).ReadString('\n')
	if err2 != nil {
		fmt.Println("1111111")
	}
	n, err3 := conn.Write([]byte(content))
	if err3 != nil {
		fmt.Println("444555666")
	}
	fmt.Printf("向服务端传过来的%d数据！！,且推出客户端", n)

}

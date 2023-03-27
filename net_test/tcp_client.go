package net

import (
	"fmt"
	"net"
)

func connect() {
	fmt.Println("客户端启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("客户端连接失败! err: ", err)
		return
	}
	fmt.Println("连接成功! conn: ", conn)

	input := "Hello World"
	n, err := conn.Write([]byte(input))
	if err != nil {
		fmt.Println("数据发送失败: err: ", err)
	}
	fmt.Printf("发送了 %d 字节数据", n)
}

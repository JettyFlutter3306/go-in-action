package net

import (
	"fmt"
	"net"
)

func handleConnection() {
	fmt.Println("服务器启动了...")
	listen, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("监听失败! err: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端等待失败! err: ", err)
		} else {
			fmt.Printf("等待连接成功! conn = %v, IP & Port: %v \n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		content := string(buffer[:n])
		if content == "exit;" {
			return
		} else {
			echo := "echo: " + content
			_, err2 := conn.Write([]byte(echo))
			if err2 != nil {
				return
			}
		}
		fmt.Printf("IP & Port: %s, content: %s", conn.RemoteAddr(), content)
	}
}

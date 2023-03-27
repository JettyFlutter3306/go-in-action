package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("客户端启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("客户端连接失败! err: ", err)
		return
	}
	fmt.Println("连接成功! conn: ", conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入失败: err: ", err)
		}
		if strings.Trim(input, "\n") == "exit;" {
			fmt.Println("client closed")
			conn.Close()
			break
		}

		n, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("数据发送失败: err: ", err)
		}
		fmt.Printf("发送了 %d 字节数据\n", n)

		buffer := make([]byte, 1024)
		n1, err := conn.Read(buffer)
		response := string(buffer[:n1])
		fmt.Printf("Server: %s, %s", conn.RemoteAddr(), response)
		fmt.Printf("读取了 %d 字节数据\n", n1)
	}
}

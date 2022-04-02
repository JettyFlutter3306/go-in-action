package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileOpen(t *testing.T) {
	// 打开文件
	path := "E:/GoProjects/Go-In-Action/resources/Test.txt"
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("文件打开出错", err)
	}

	// 输出文件
	fmt.Printf("文件: %v", file)

	// 关闭文件
	e := file.Close()
	if err != nil {
		fmt.Println("文件关闭失败", e)
	}
}

func TestFileIO(t *testing.T) {
	path := "E:/GoProjects/Go-In-Action/resources/Test.txt"

	// 返回内容是 []byte, err
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("读取出错, 错误为: ", err)
	}

	// 读取成功
	fmt.Printf("%v", string(content))
}

func TestBufferInput(t *testing.T) {
	// 打开文件
	path := "E:/GoProjects/Go-In-Action/resources/Test.txt"
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("文件打开出错", err)
	}

	// 当函数退出时,关闭文件,防止内存溢出
	defer file.Close()

	// 创建一个缓冲流
	reader := bufio.NewReader(file)

	// 读取操作
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		print(str)
	}

	println("文件读取完成...")
}

func TestFileOutput(t *testing.T) {
	// 写入文件操作
	path := "E:/GoProjects/Go-In-Action/resources/Test.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		println("打开文件失败", err)
		return
	}

	// 及时将文件关闭
	defer file.Close()

	// 写入文件操作 --> 使用IO流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString("你好世界, 我是最牛逼的程序员!\n")
	}

	// 流带缓冲区,刷新数据
	_ = writer.Flush()
}

func TestFileCopy(t *testing.T) {
	src := "E:/GoProjects/Go-In-Action/resources/Test.txt"
	des := "E:/Test.txt"

	content, err := ioutil.ReadFile(src)
	if err != nil {
		println("文件读取有问题")
		return
	}

	err1 := ioutil.WriteFile(des, content, 0666)

	if err1 != nil {
		println("文件输出有误!")
	}

}

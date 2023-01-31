package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 定义管道,声明管道
func TestChannel(t *testing.T) {
	var channel1 = make(chan int, 3)
	fmt.Printf("channel的值: %v\n", channel1)

	// 向管道存放数据
	var values = []int{2, 5, 9}
	for _, value := range values {
		channel1 <- value
	}

	fmt.Printf("管道的实际长度: %v, 管道的容量: %v\n", len(channel1), cap(channel1))
}

func TestClose(t *testing.T) {
	var channel1 = make(chan int, 3)
	var values = []int{2, 5, 9}

	for _, value := range values {
		channel1 <- value
	}

	close(channel1)

	// 管道关闭之后,读取数据还是可以的
	num := <-channel1
	fmt.Println(num)
}

func TestTraverseChannel(t *testing.T) {
	var channel = make(chan int, 100)
	for i := 0; i < 100; i++ {
		channel <- i
	}

	// 在遍历前,如果没有关闭管道,就会出现deadlock
	close(channel)
	for v := range channel {
		fmt.Println("value = ", v)
	}
}

func writeData(channel chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		channel <- i
		fmt.Println("写入的数据为: ", i)
		time.Sleep(time.Second)
	}
	close(channel)
}

func readData(channel chan int) {
	defer wg.Done()
	for v := range channel {
		fmt.Println("读取的数据为: ", v)
		time.Sleep(time.Second)
	}
}

var wg sync.WaitGroup

// 模拟生产者消费者模型
func TestRWChannel(t *testing.T) {
	var channel = make(chan int, 50)
	wg.Add(2)
	go writeData(channel)
	go readData(channel)
	wg.Wait()
}

func TestRWOnly(t *testing.T) {
	// 声明只写管道
	var channel chan<- int = make(chan int, 3)
	channel <- 20
	fmt.Println("channel: ", channel)

	// 声明只读管道
	var channel1 <-chan int = make(chan int, 3)
	fmt.Println("channel1: ", channel1)
}

func TestSelectChannel(t *testing.T) {
	var chan1 = make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		chan1 <- 10
	}()

	var chan2 = make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "string channel"
	}()

	// 取数据就是阻塞操作
	select {
	case v := <-chan1:
		fmt.Println("channel1: ", v)
	case v := <-chan2:
		fmt.Println("channel2: ", v)
	default:
		fmt.Println("防止Select阻塞")
	}
}

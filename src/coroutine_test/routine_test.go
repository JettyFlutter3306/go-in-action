package thread_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func printGolang() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func TestRoutine(t *testing.T) {
	// 开启一个协程
	go printGolang()
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 使用匿名函数启动协程
func TestAnonymousRoutine(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

var group sync.WaitGroup

// 使用WaitGroup控制协程退出
func TestWaitGroup(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 协程开启加1
		group.Add(1)
		go func(n int) {
			// 协程执行完成减1,放在最后执行
			defer group.Done()
			fmt.Println(n)
		}(i)
	}

	// 使主线程一直阻塞,WaitGroup都减为0了,就停止
	// 可以类比 Java并发包的 CountDownLatch
	group.Wait()
}

var total int
var lock sync.Mutex

// 交叉执行最后的结果不一定是0
func TestSharedVariable(t *testing.T) {
	group.Add(2)
	go func() {
		defer group.Done()
		for i := 0; i < 10000; i++ {
			lock.Lock()
			total += 1
			lock.Unlock()
		}
	}()

	go func() {
		defer group.Done()
		for i := 0; i < 10000; i++ {
			lock.Lock()
			total -= 1
			lock.Unlock()
		}
	}()

	group.Wait()
	fmt.Println(total)
}

package concurrent

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

var wg sync.WaitGroup

func PrintNumber() {
	wg.Add(2)
	// 定义打印奇数协程
	printOdd := func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	// 定义打印偶数协程
	printEven := func() {
		defer wg.Done()
		for i := 0; i < 10; i += 2 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	go printOdd()
	go printEven()
	wg.Wait()
}

// 同时启动多个goroutine输出不同的数字
func GoroutineRandom() {
	workerNum := 10
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}

func GoroutineNum() {
	// 统计当前存在的goroutine的数量
	go func() {
		for {
			fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 启动大量的goroutine
	for {
		go func() {
			time.Sleep(100 * time.Second)
		}()
	}
}

// 协程池控制并发量
func GoroutinePool() {
	go func() {
		for {
			fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 初始化协程池
	size := 1024
	pool, err := ants.NewPool(size)
	if err != nil {
		log.Fatal(err)
	}
	// 函数返回之前关闭协程池
	defer pool.Release()

	// 利用pool调用需要并发的大量goroutine
	for {
		// 向pool中提交一个执行的goroutine
		err := pool.Submit(func() {
			_ = make([]int, 1024)
			fmt.Println("in goroutine")
			time.Sleep(time.Second * 100)
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}

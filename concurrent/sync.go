package concurrent

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func SyncErr() {
	wg := sync.WaitGroup{}
	counter := 0
	gs := 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// ++ 操作不是原子操作
				// 1. 获取当前的counter变量
				// 2. +1
				// 3. 赋值新值到counter
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter: ", counter)
}

func SyncLock() {
	wg := sync.WaitGroup{}
	counter := 0
	gs := 100

	wg.Add(gs)
	// 创建互斥锁
	lock := sync.Mutex{}
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// ++ 操作不是原子操作
				// 1. 获取当前的counter变量
				// 2. +1
				// 3. 赋值新值到counter
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter: ", counter)
}

func SyncMutex() {
	wg := sync.WaitGroup{}
	var lock sync.Mutex
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("before lock: ", n)
			lock.Lock()
			fmt.Println("locked: ", n)
			time.Sleep(time.Second)
			lock.Unlock()
			fmt.Println("after lock: ", n)
		}(i)
	}
	wg.Wait()
}

func SyncLockAndNo() {
	wg := sync.WaitGroup{}
	counter := 0
	gs := 100

	wg.Add(gs)
	// 创建互斥锁
	lock := sync.Mutex{}
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// ++ 操作不是原子操作
				// 1. 获取当前的counter变量
				// 2. +1
				// 3. 赋值新值到counter
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}

	wg.Add(1)
	// 和上面不是一把锁，结果会有错误
	var lock1 sync.Mutex
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			lock1.Lock()
			counter++
			lock1.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("counter: ", counter)
}

func SyncRLock() {
	wg := sync.WaitGroup{}
	var rwLock sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// rwLock.Lock()
			rwLock.RLock()
			fmt.Println(time.Now())
			time.Sleep(time.Second)
			rwLock.RUnlock()
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		rwLock.RLock()
		fmt.Println(time.Now(), "Lock")
		time.Sleep(time.Second)
		rwLock.RUnlock()
	}()
	wg.Wait()
}

func SyncMapErr() {
	m := map[string]int{}
	// 写操作
	go func() {
		for {
			m["key"] = 0
		}
	}()

	// 读操作
	go func() {
		for {
			_ = m["key"]
		}
	}()

	select {}
}

func SyncMap() {
	var m sync.Map
	go func() {
		for {
			m.Store("key", 0)
		}
	}()

	go func() {
		for {
			_, _ = m.Load("key")
		}
	}()

	select {}
}

func SyncMapMethod() {
	wg := sync.WaitGroup{}
	var m sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// 存储
			m.Store(n, fmt.Sprintf("value(%d)", n))
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(m.Load(n))
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	m.Delete(4)
}

func SyncPool() {
	// 原子计数器
	var counter int32 = 0

	// 定义元素的Newer，创建器
	elFactory := func() any {
		atomic.AddInt32(&counter, 1)
		return new(bytes.Buffer)
	}

	// Pool初始化
	// pool := sync.Pool{
	// 	New: elFactory,
	// }

	// 并发申请和交回元素
	workerNum := 1024 * 1024
	wg := sync.WaitGroup{}
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			defer wg.Done()
			// 申请元素，通常需要断言为特定类型
			// buffer := pool.Get().(*bytes.Buffer)
			buffer := elFactory().(*bytes.Buffer)
			// defer pool.Put(buffer)
			// 使用元素
			_ = buffer.String()
		}()
	}

	wg.Wait()
	fmt.Println("number is : ", counter)
}

// sync.Once 只执行一次
func SyncOnce() {
	// 初始化config变量
	config := make(map[string]string)

	// 初始化sync.Once
	var once sync.Once

	// 加载配置函数
	loadConfig := func() {
		// 利用once.Do来执行
		once.Do(func() {
			config = map[string]string{
				"varInt": fmt.Sprintf("%d", rand.Int31()),
			}
			fmt.Println("config loaded.")
		})
	}

	// 模拟多个goroutine，多次调用加载配置
	// 测试加载配置操作，执行了几次
	workers := 10
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			loadConfig()
			// 使用配置
			_ = config
		}()
	}

	wg.Wait()
}

func SyncCond() {
	var wg sync.WaitGroup
	var data []int
	dataLen := 1024 * 1024

	// 创建sync.Cond
	condition := sync.NewCond(&sync.Mutex{})

	// 结束数据的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		condition.L.Lock()
		defer condition.L.Unlock()
		for i := 0; i < dataLen; i++ {
			data = append(data, i*i)
		}
		time.Sleep(time.Second * 2)

		// 广播，可选的需要锁定
		condition.Broadcast()
		fmt.Println("condition broadcast")
	}()

	// 处理数据的goroutine
	const workers = 8
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			// 在数据接收完之前等待
			condition.L.Lock()
			// if len(data) < dataLen {
			// 	condition.Wait()
			// }
			for len(data) < dataLen {
				condition.Wait()
			}
			fmt.Println("处理数据, 数据长度: ", len(data))
			condition.L.Unlock()
		}()
	}
	wg.Wait()
}

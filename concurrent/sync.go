package concurrent

import (
	"fmt"
	"sync"
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

package thread

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 加入读写锁
var rwLock sync.RWMutex
var group1 sync.WaitGroup

func read() {
	defer group1.Done()
	rwLock.RLock()
	fmt.Println("开始读取数据...")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功...")
	rwLock.RUnlock()
}

func write() {
	defer group1.Done()
	rwLock.Lock()
	fmt.Println("开始修改数据...")
	time.Sleep(time.Second * 5)
	fmt.Println("修改数据成功...")
	rwLock.Unlock()
}

func TestRWLock(t *testing.T) {
	group1.Add(6)
	// 启动协程 ---> 读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	group1.Wait()
}

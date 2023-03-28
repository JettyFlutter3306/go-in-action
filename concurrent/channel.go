package concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func ChannelOps() {
	// 初始化，无缓冲channel
	ch := make(chan int)

	// send
	go func() {
		ch <- 10
	}()

	// receive
	go func() {
		v := <-ch
		fmt.Println("received value: ", v)
	}()

	time.Sleep(time.Second)

	// close
	close(ch)
}

func ChannelFor() {
	wg.Add(1)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range ch {
			fmt.Println("received from ch, element is : ", e)
		}
	}()

	wg.Wait()
}

func ChannelSync() {
	wg.Add(2)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
			println("Sent ", i, "\tNow: ", time.Now().Format("15.04.05.9999999"))
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for e := range ch {
			println("Received ", e, "\tNow: ", time.Now().Format("15.04.05.9999999"))
			time.Sleep(time.Second * 3)
		}
	}()

	wg.Wait()
}

func ChannelAsync() {
	wg.Add(2)
	// 加一个缓冲区
	ch := make(chan int, 10)
	println("Cap: ", cap(ch))

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
			println("Sent ", i, "\tNow: ", time.Now().Format("15.04.05.9999999"), ", Len: ", len(ch))
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for e := range ch {
			println("Received ", e, "\tNow: ", time.Now().Format("15.04.05.9999999"), ", Len: ", len(ch))
			time.Sleep(time.Second * 3)
		}
	}()

	wg.Wait()
}

func ChannelGoroutineControl() {
	// 独立的goroutine输出数量
	go func() {
		for {
			println("goroutine number: ", runtime.NumGoroutine())
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// 初始化channel，设置缓冲大小
	const size = 1024
	ch := make(chan struct{}, size)

	// 并发的goroutine
	for {
		// 管道已满就阻塞
		ch <- struct{}{}
		go func() {
			time.Sleep(time.Second * 10)
			// 移除管道内的元素
			<-ch
		}()
	}
}

func ChannelDirectional() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 10)

	wg.Add(2)
	// 使用双向channel为单向channel赋值
	go getElement(ch, wg)
	go setElement(ch, 100, wg)

	wg.Wait()
}

func getElement(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Received: ", <-ch)
}

func setElement(ch chan<- int, v int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- v
	println("Sent: ", v)
}

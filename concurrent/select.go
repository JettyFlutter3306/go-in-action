package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SelectStmt() {
	var a [4]int
	var c1, c2, c3, c4 = make(chan int), make(chan int), make(chan int), make(chan int)
	var i1, i2 = 0, 42

	go func() {
		c1 <- 10
	}()

	go func() {
		<-c2
	}()

	go func() {
		close(c3)
	}()

	go func() {
		c4 <- 100
	}()

	// select多路监听goroutine
	go func() {
		select {
		case i1 = <-c1:
			println("received from c1: ", i1)
		case c2 <- i2:
			println("sent ", i2, " to channel c2")
		case i3, ok := <-c3:
			if ok {
				println("received from channel c3: ", i3)
			} else {
				println("channel c3 was closed.")
			}
		case a[f()] = <-c4:
			println("received from channel c4: ", a[f()])
		default:
			println("no channel operations.")
		}
	}()
}

func f() int {
	print("f() is running...")
	return 2
}

func SelectFor() {
	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			select {
			case v := <-ch:
				println("received from channel: ", v)
			}
		}
	}()

	time.Sleep(time.Second * 3)
}

func SelectBlock() {
	println("before select")
	select {}
	println("after select")
}

func SelectNilChannel() {
	ch := make(chan int)

	go func() {
		rand.Seed(time.Now().Unix())
		for {
			ch <- rand.Intn(10)
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		// 设置定时器
		t := time.After(time.Second * 3)
		sum := 0
		for {
			select {
			case v := <-ch:
				println("received value: ", v)
				sum += v
			case <-t:
				ch = nil
				println("ch was set nil, sum is ", sum)
			}
		}
	}()

	time.Sleep(time.Second * 5)
}

func SelectNonBlock() {
	counter := 10
	max := 20
	rand.Seed(time.Now().UnixMilli())
	answer := rand.Intn(max) // [0, 19]
	println("The answer is ", answer)
	println("========================")

	// 模拟猜数
	bingoCh := make(chan int, counter)
	wg := sync.WaitGroup{}
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		// 每一个协程代表一个猜数的人
		go func() {
			defer wg.Done()
			result := rand.Intn(max)
			println("someone guessed: ", result)
			if result == answer {
				bingoCh <- result
			}
		}()
	}
	wg.Wait()

	println("===========================")
	select {
	case result := <-bingoCh:
		println("someone hit the answer: ", result)
	default:
		// 非阻塞保证
		println("no one hit the answer.")
	}
}

func SelectRace() {
	// 模拟查询结果，需要与具体的query建立联系
	// 用于通信的channel，数据与停止信号
	type Rows struct {
		Index int
	}
	const QueryNum = 8
	ch := make(chan Rows, 1)
	stopChs := [QueryNum]chan struct{}{}
	for i := range stopChs {
		stopChs[i] = make(chan struct{})
	}
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixMilli())

	// 模拟query查询，每个查询持续不同的时间
	wg.Add(QueryNum)
	for i := 0; i < QueryNum; i++ {
		// 每个query
		go func(i int) {
			defer wg.Done()
			// 模拟执行时间
			randD := rand.Intn(1000)
			println("querier ", i, " start fetch data, need duration is ", randD, " ms.")

			// 查询结果的channel
			chRst := make(chan Rows, 1)

			// 执行查询工作
			go func() {
				// 模拟时长
				time.Sleep(time.Duration(randD) * time.Millisecond)
				chRst <- Rows{
					Index: i,
				}
			}()

			// 监听查询结果和停止信号channel
			select {
			case rows := <-chRst:
				println("querier ", i, " gets result.")
				if len(ch) == 0 {
					ch <- rows
				}
			case <-stopChs[i]:
				println("querier ", i, " is stopping.")
				return
			}
		}(i)
	}

	// 等待第一个查询结果的反馈
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// 等待第一个查询结果
		case rows := <-ch:
			println("get first result from querier ", rows.Index, " stop other queriers")
			for i := range stopChs {
				// 当前返回结果的goroutine不需要了，因为已经结束了
				if i == rows.Index {
					continue
				}
				// 往结束信号的管道里面填入数据，表示协程运行结束
				stopChs[i] <- struct{}{}
			}
		// 计划一个超时时间
		case <-time.After(time.Second * 5):
			println("all queriers timeout.")
			for i := range stopChs {
				stopChs[i] <- struct{}{}
			}
		}
	}()

	wg.Wait()
}

func SelectAll() {
	type Content struct {
		Subject string
		Tags    []string
		Views   int
		// 标识操作了资源哪个部分
		part string
	}

	const (
		PartSubject = "subject"
		PartTags    = "tags"
		PartViews   = "views"
	)

	// 同步停止的信号的map
	stopChs := map[string]chan struct{}{
		PartSubject: make(chan struct{}),
		PartTags:    make(chan struct{}),
		PartViews:   make(chan struct{}),
	}

	// 接收和发送操作的通信channel
	ch := make(chan Content, len(stopChs))
	wg := sync.WaitGroup{}
	to := time.After(time.Millisecond * 800)
	rand.Seed(time.Now().UnixMilli())

	// goroutine执行每个部分的获取
	for part := range stopChs {
		wg.Add(1)
		go func(p string) {
			// 初始化
			defer wg.Done()
			// 模拟延时获取资源
			randD := rand.Intn(1000)
			println("start fetching ", p, " data, need duration ", randD, " ms.")
			chRst := make(chan Content, 1)

			go func() {
				time.Sleep(time.Duration(randD) * time.Millisecond)
				content := Content{
					part: p,
				}
				switch p {
				case PartSubject:
					content.Subject = "Subject of content"
				case PartTags:
					content.Tags = []string{"go", "Goroutine", "Channel", "select"}
				case PartViews:
					content.Views = 1024
				}
				// 发送到chRst
				chRst <- content
			}()

			// 监控资源获取成功还是超时到达
			select {
			// 查询到结果
			case rst := <-chRst:
				println("querier ", p, " gets result.")
				ch <- rst
			// 超时到了
			case <-stopChs[p]:
				println("querier ", p, " is stopping.")
				return
			}
		}(part)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 初始化资源
		content := Content{}
		received := map[string]struct{}{}

		// 等待接收或超时到期
	FLAG:
		for {
			select {
			// 超时
			case <-to:
				println("querier timeout. Content is incomplete.")
				// 超时时间到要通知未完成的goroutine结束
				// 遍历stopCh，判定是否存在于received中即可
				for part := range stopChs {
					if _, ok := received[part]; !ok {
						stopChs[part] <- struct{}{}
					}
				}
				close(ch)
				// 不再继续监听，结束循环
				break FLAG
			// 有处理的业务
			case rst := <-ch:
				println("received some part ", rst.part)
				switch rst.part {
				case PartSubject:
					content.Subject = rst.Subject
					received[PartSubject] = struct{}{}
				case PartTags:
					content.Tags = rst.Tags
					received[PartTags] = struct{}{}
				case PartViews:
					content.Views = rst.Views
					received[PartViews] = struct{}{}
				}
				// 完成标志
				finished := true
				// 确认是否都接收了
				for part := range stopChs {
					if _, ok := received[part]; !ok {
						finished = false
						break
					}
				}
				// 判定finished，说明全部已经处理完毕了
				if finished {
					println("all queriers finished. Content is complete.")
					close(ch)
					break FLAG
				}
			}
		}
		fmt.Println("Content: ", content)
	}()
	wg.Wait()
}

func SelectChannelCloseSignal() {
	wg := sync.WaitGroup{}
	// 定义无缓冲channel
	ch := make(chan struct{})

	// 用来close管道
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		fmt.Println("Emit the signal, close(ch)")
		close(ch)
	}()

	// 接收ch，表示接受信号
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("收到信号, <-ch")
				return
			default:
				// 正常的业务逻辑
				fmt.Println("业务逻辑处理中...")
				time.Sleep(time.Millisecond * 300)
			}
		}
	}()
	wg.Wait()
}

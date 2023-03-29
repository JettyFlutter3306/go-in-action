package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func TimerA() {
	t := time.NewTimer(time.Second)
	fmt.Printf("timer is set, %s\n", time.Now())

	now := <-t.C
	fmt.Printf("time is up, %s\n", now)
}

func TimerB() {
	ticker := time.NewTicker(time.Millisecond * 400)
	ch := make(chan int)
	go func() {
		for _ = range ticker.C {
			ch <- rand.Intn(10)
		}
	}()

	// 定时器控制每局时间
	t := time.NewTimer(time.Second * 3)
	counter, hit, miss, result := 5, 0, 0, 4
	for i := 0; i < counter; i++ {
	FLAG:
		for {
			select {
			case v := <-ch:
				fmt.Println("You guess is ", v)
				if result == v {
					fmt.Println("Bingo. You hit the answer.")
					hit++
					t.Reset(time.Second * 3)
					break FLAG
				}
			case <-t.C:
				fmt.Println("Time is up, miss the target.")
				miss++
				// 重新创建定时器，进行下一轮
				t = time.NewTimer(time.Second * 3)
				break FLAG
			}
		}
	}
	ticker.Stop()
	fmt.Printf("Game Over, hit: %d, miss: %d\n", hit, miss)
}

func TickerA() {
	ticker := time.NewTicker(time.Second)
	timer := time.After(time.Second * 5)

	// 周期性执行
FLAG:
	for t := range ticker.C {
		fmt.Printf("HeartBeat, http.Get(http://domain), time: %s\n", t)
		select {
		case <-timer:
			fmt.Println("Heartbeat finished.")
			ticker.Stop()
			break FLAG
		default:
		}
	}
}

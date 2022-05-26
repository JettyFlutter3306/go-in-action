package thread

import (
	"fmt"
	"strconv"
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

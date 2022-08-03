package process_test

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	/*
		定义一个功能实现 1+2+3+4+5
	*/

	var ans = 0

	/*
		Caution: 在for循环中不可以使用 var 声明一个变量 !!
				 需要使用 := 符号
		Golang中只有for循环,没有 while, do-while 循环
	*/
	for i := 1; i <= 5; i++ {
		ans += i
	}

	fmt.Println(ans)

	var s1 = "Hello Golang"

	//使用普通for循环
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c", s1[i]) //打印字节的ASCII码值
	}

	fmt.Println()

	/*
		使用 for range 遍历
	*/
	for i, value := range s1 {
		fmt.Printf("索引: %d,内容: %c \n", i, value)
	}
}

package process_test

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	// 实现功能,如果口罩的库存小于30个,提示库存不足
	// var count = 20

	/*
		在Golang中,if后面可以并列地加入变量的定义
	*/
	// if count < 30 {
	//	fmt.Println("库存不足!")
	// }

	if count := 50; count < 30 {
		fmt.Println("库存不足!")
	} else {
		fmt.Println("库存足够!")
	}

	var score = 60

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

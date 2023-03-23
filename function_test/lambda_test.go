package main

import (
	"fmt"
	"testing"
)

func TestLambda(t *testing.T) {
	// 定义一个匿名函数,并调用
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println(result)

	/*
		将匿名函数赋值给一个变量,这个变量实际上就是函数类型的变量
		sub等价于匿名函数
	*/
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println(sub(90, 88))
}

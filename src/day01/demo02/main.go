package main

import (
	"fmt"
	"unsafe"
)

/*
GO语言中,假如变量只声明而不使用则会报错
*/
func main() {

	var num1 int8 = 120 //定义一个整数类型
	fmt.Println(num1)

	var num2 uint8 = 200
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的类型是: %T", num3)
	fmt.Println(unsafe.Sizeof(num3))

	/**
	在GO语言中,要尽可能的使用占用内存空间的小的数据类型
	表示学生年龄,使用 uint 关键字等价于 byte
	*/
	var age uint8 = 18
	fmt.Print(age)

}

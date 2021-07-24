package main

import "fmt"

func main() {

	/*
		定义字符类型的数据
		本质上就是一个整数,所以可以直接参与运算,参照ASCII码表
		Golang使用的是UTF-8编码方式
		UTF-8是Unicode的一种编码方案
	*/
	var c1 byte = 'a'
	fmt.Println(c1) //97

	var c2 byte = '6'
	fmt.Println(c2) //54

	var c3 byte = '('
	fmt.Println(c3) //40

	var c4 = '中'
	fmt.Println(c4) //对应的是Unicode编码表

	var c5 = 'A'
	fmt.Printf("c5对应的字符是: %c", c5)

}

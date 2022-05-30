package base_test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	var age = 18
	fmt.Println(&age) //打印地址: 0xc00000a098

	/*
		定义一个指针变量
		p指针变量的名字
		p对应的类型是: *int,指针类型(可以理解为指向 int 的指针)
		&age就是一个地址,是p变量的具体的值
	*/
	var p = &age
	fmt.Println(p)
	fmt.Println("p本身这个存储空间的地址为: ", &p)

	// 想获取指针指向的区域的数据
	fmt.Println("p指向的内存区域的数值为: ", *p)
}

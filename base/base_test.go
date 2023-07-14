package base_test

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
定义在函数外的变量是全局变量
*/
var n7 = 100
var n8 = 9.7

/*
一次性声明多个全局变量
*/
var (
	n9  = 500
	n10 = "洛必达"
)

/*
定义在函数中的变量是局部变量
*/
func TestStateDate(t *testing.T) {
	// 第一种变量使用的方式
	var num int = 18
	fmt.Println(num)

	// 第二种,指定变量类型但是没有赋值
	var num1 int
	fmt.Println(num1) // 默认是0

	// 第三种,自动类型推断
	var num3 = 10
	fmt.Println(num3)

	// 第四种,省略var关键字 使用  :=  符号赋值
	num4 := num3
	fmt.Println(num4)

	fmt.Println("------------------------------------------------------")

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, name, n5 = 10, "jack", 7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	n6, height := 6.9, 100.6
	fmt.Println(n6)
	fmt.Println(height)

	fmt.Println("--------------------------------------------------------")

	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)

	fmt.Printf("num3的类型是: %T", num3)
	fmt.Println(unsafe.Sizeof(num3))

	/**
	在GO语言中,要尽可能的使用占用内存空间的小的数据类型
	表示学生年龄,使用 uint 关键字等价于 byte
	*/
	var age uint8 = 18
	fmt.Print(age)

	/*
		定义字符类型的数据
		本质上就是一个整数,所以可以直接参与运算,参照ASCII码表
		Golang使用的是UTF-8编码方式
		UTF-8是Unicode的一种编码方案
	*/
	var c1 byte = 'a'
	fmt.Println(c1) // 97

	var c2 byte = '6'
	fmt.Println(c2) // 54

	var c3 byte = '('
	fmt.Println(c3) // 40

	var c4 = '中'
	fmt.Println(c4) // 对应的是Unicode编码表

	var c5 = 'A'
	fmt.Printf("c5对应的字符是: %c", c5)
}

func filter(arr []int, predicate func(int) bool) []int {
	var ans []int
	for _, v := range arr {
		if predicate(v) {
			ans = append(ans, v)
		}
	}
	return ans
}

func TestType(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	// 保留偶数
	ans := filter(arr, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(ans)
}

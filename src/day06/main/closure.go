package main

import "fmt"

/*
函数功能求和,
函数的名字,getSum参数为空
返回值为一个函数,并且这个函数的参数是 int 类型,返回值也是 int 类型
*/
func getSum() func(int) int {

	var sum = 10 //sum对于返回的匿名函数可以一直使用,匿名函数和外面的变量一直保存在内存中,故可以一直使用!

	return func(num int) int {
		sum += num

		return sum
	}

}

/*
闭包的定义:  (closure)
	闭包就是一个函数和其他相关的引用环境组合的一个整体!!
	返回的匿名函数 + 匿名函数以外的变量 num

	闭包就是能够读取其他函数内部变量的函数。
	例如在javascript中，只有函数内部的子函数才能读取局部变量，
	所以闭包可以理解成“定义在一个函数内部的函数“。
	在本质上，闭包是将函数内部和函数外部连接起来的桥梁。

	不使用闭包的时候,我想保留的值,不可以反复使用
	闭包使用的场景: 闭包可以保留上次引用的某个值,我们传入一次的值就可以反复使用了
*/
func main() {

	fun1 := getSum() //这里返回一个函数对象

	fmt.Println(fun1(10)) //20
	fmt.Println(fun1(10)) //30
	fmt.Println(fun1(10)) //40
	fmt.Println(fun1(10)) //50

}

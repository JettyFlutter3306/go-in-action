package main

import (
	"fmt"
	"testing"
)

/*
六十一讲
学习使用函数,如果返回值就只有一个的话,那么()可以省略不写
func name(v1, v2, ...) (return value) {

	return (return value)
}

函数名称开头大写表示: public 权限 在别的包内也可使用
函数名称开头小写表示: private 权限 在别的包内不可访问,只可以在本包内使用
返回可以用 _ 符号忽略
Golang中,不支持  "函数重载" !!
*/
func TestFunction(t *testing.T) {
	//i := sum(10, 20)
	//
	//fmt.Println(i)
	//
	//sum, difference := sumAndDifference(20, 30)
	//
	//fmt.Println(sum)
	//fmt.Println(difference)
	//
	//fmt.Println(fun1(2, 3, 4, 5, 9, 10))
	num := 10
	//fun2(&num, 100)

	fun3(&num, 200, fun2)
	fmt.Println(num)
}

func sum(n1 int, n2 int) int {

	return n1 + n2
}

func sumAndDifference(n1 int, n2 int) (int, int) {

	sum := n1 + n2
	difference := n1 - n2

	return sum, difference
}

/*
函数内部处理可变参数时,一律当成切片数据类型处理
基本数据类型和数组是按照值传递的,数组不是引用传递!
*/
func fun1(args ...int) int {

	var sum = 0

	for i := range args {
		sum += i
	}

	return sum
}

/*
通过引用传递从而修改变量的值
参数类型为指针
*/
func fun2(n1 *int, n2 int) {

	*n1 = n2
}

/*
函数作为参数传递  闭包!
*/
func fun3(num1 *int, num2 int, funTest func(*int, int)) {
	funTest(num1, num2)
}

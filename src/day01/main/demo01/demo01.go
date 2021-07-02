package main

import "fmt"

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

//第十八讲
func main() {

	/*
		定义在函数中的变量是局部变量
	*/

	//第一种变量使用的方式
	var num int = 18
	fmt.Println(num)

	//第二种,指定变量类型但是没有赋值
	var num1 int
	fmt.Println(num1) //默认是0

	//第三种,自动类型推断
	var num3 = 10
	fmt.Println(num3)

	//第四种,省略var关键字 使用  :=  符号赋值
	num4 := num3
	fmt.Println(num4)

	fmt.Println("------------------------------------------------------")

	//声明多个变量
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

}

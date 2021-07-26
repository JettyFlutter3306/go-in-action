package main

import "fmt"

func main() {

	/*
		+加号
		1.整数	2.相加操作	3.字符串拼接
	*/
	var n1 = +10
	fmt.Println(n1)

	var n2 = 4 + 7
	fmt.Println(n2)

	var s1 = "abc" + "def"
	fmt.Println(s1)

	/*
		/除号
	*/
	fmt.Println(10 / 3)   //两个int类型数据运算,结果一定是整数类型
	fmt.Println(10.0 / 3) //浮点数参与运算,结果是浮点数

	/*
		%取模符号
	*/
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)

	/*
		++操作
		Golang中 ++ -- 操作非常简单,但是只能单独使用,不能参与到运算中去
		Golang中 ++ -- 只能写在变量后面,不可以写在变量前面 --a ++a 是错误写法
	*/
	var a = 10
	a++
	fmt.Println(a)
}

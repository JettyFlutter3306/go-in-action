package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n1 int64 = 10
	s1 := strconv.FormatInt(n1, 10)

	fmt.Printf("s1 = %q \n", s1)

	/*
		第二个参数, 'f' (-ddd.dddd)
		第三个参数, 保留小数点后面的位数
	*/
	var n2 float64 = 4.29
	s2 := strconv.FormatFloat(n2, 'f', 9, 64)
	fmt.Printf("s2 = %q \n", s2)

	var s3 = "true"
	var b bool

	/*
		ParseBool这个函数的返回值有两个: (value bool, err error)
		value就是我们得到的bool类型数据
		err指的是错误
		如果我们只关注得到的布尔类型的数据,那么err就可以忽略  使用 _ 符号
	*/
	b, _ = strconv.ParseBool(s3)
	fmt.Printf("b的类型是: %T, b = %v \n", b, b)
}

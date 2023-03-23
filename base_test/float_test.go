package base_test

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	/*
		float可以表示正的浮点数,也可以表示负的浮点数
		浮点数可以使用十进制表示形式,也可以使用科学计数法表示形式,使用 E (e) 都可以表示10的幂次

		浮点数可能会有精度损失,所以通常情况下,推荐使用 float64 数据类型
		GO语言中浮点数默认类型是 float64
	*/
	// 定义浮点类型
	var num1 float32 = 3.14
	fmt.Println(num1)

	var num2 float32 = -3.14
	fmt.Println(num2)

	var num3 float32 = 314e-2
	fmt.Println(num3)

	var num4 float32 = 314e2
	fmt.Println(num4)

	var num5 float32 = 314e2
	fmt.Println(num5)

	var num6 float32 = 256.000000916
	fmt.Println(num6)

	var num7 float64 = 256.000000916
	fmt.Println(num7)

	var num8 = 256.0008078
	fmt.Printf("num8的数据类型是: %T", num8)
}

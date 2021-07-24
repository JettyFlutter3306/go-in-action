package main

import "fmt"

func main() {

	/*
		定义一个字符串
		字符串不可变: 指的是字符串一旦定义之后,其中字符的值就不可变
		字符串的表示形式:
			(1)字符串中没有特殊字符,字符串的表示形式用双引号
			(2)如果有特殊字符,那么就用反引号 ``
	*/
	var s1 string = "抓住机遇,拥抱未来! Golang!"
	fmt.Println(s1)

	var s2 = `
		package main

		import "fmt"

		func main() {
		var s1 string = "抓住机遇,拥抱未来! Golang!"
		fmt.Println(s1)
	`
	fmt.Println(s2)

	/*
		在这里隐式自动数据类型转换不生效,必须得手动的显示转换!!
	*/
	var n1 int = 100
	var n2 float32 = float32(n1)
	fmt.Println(n1)
	fmt.Println(n2)

}

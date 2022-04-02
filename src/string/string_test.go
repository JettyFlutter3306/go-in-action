package string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

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

	/*
		使用内置函数:
			1.统计字符串的长度  len(str)
			2.遍历字符串  for ... range
			3.字符串转数字  strconv.Atoi(str)
			4.数字转字符串  strcoonv.Itoa(number)
			5.统计一个字符串有几个指定的子串  strings.Count()
			6.不区分大小写的字符串的比较  strings.EqualFold()
			7.返回子串在字符串第一次出现的索引值,如果没有则返回 -1  strings.Index()
			8.字符串的替换  strings.Replace()
			9.分割字符串  strings.Replace()
			10.大小写转换 strings.ToLower()  strings.ToUpper()
			11.去除字符串两边的空格  strings.TrimSpace()
			12.去除字符串两边指定的字符  strings.Trim()
			13.去除字符串(左边)右边的指定的字符  strings.TrimLeft()  stings.TrimRight()
			14.判断字符串是否以指定的字符串打头(结尾)  strings.HasPrefix()  strings.HasSuffix()

	*/
	var str = "hello golang  Go语言" //在golang中,汉字是utf-8字符集,一个汉字3个字节

	fmt.Println(len(str)) //22

	//遍历方式一
	//for i, value := range str {
	//	fmt.Printf("索引为: %d,具体的值为: %c \n", i, value)
	//}

	//遍历方式二  切片
	//r:= []rune(str)
	//
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("%c \n", r[i])  //r[i]表示对应的ASCII码
	//}

	//字符串转为数字
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	//数字转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	count := strings.Count("galangandjava", "ga")
	fmt.Println(count)

	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	fmt.Println("hello" == "HELLO") //区分大写小比较

	index := strings.Index("golangandjava", "ga")
	fmt.Println(index)

	replace := strings.Replace("golangjavagogo", "go", "golang", -1)
	fmt.Println(replace)

	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

}

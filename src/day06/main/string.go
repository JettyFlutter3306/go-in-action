package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

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
			15.

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

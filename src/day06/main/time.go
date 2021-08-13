package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		时间和日期的函数,需要导入time包,要获取当前的时间,那么就调用time包的Now()函数
		默认格式: 2021-08-13 16:24:35.5352656 +0800 CST m=+0.006611301
		Now()函数返回的类型是一个结构体 类型是: time.Time
	*/
	now := time.Now()
	fmt.Println(now)

	//调用结构体中的方法
	fmt.Printf("年: %v \n", now.Year())
	fmt.Printf("月: %v \n", int(now.Month()))
	fmt.Printf("日: %v \n", now.Day())
	fmt.Printf("时: %v \n", now.Hour())
	fmt.Printf("分: %v \n", now.Minute())
	fmt.Printf("秒: %v \n", now.Second())

	fmt.Println("----------------------------------------")

	fmt.Printf("当前年月日: %d-%d-%d 时分秒: %d:%d:%d \n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	dateStr := now.Format("2006-01-02 15:04:05")
	fmt.Println(dateStr)

}

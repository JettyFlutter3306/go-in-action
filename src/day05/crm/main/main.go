package main //使用package进行包的声明,建议: 包的声明这个包和所在的文件夹同名
//main包是程序的入口包,一般main函数会放在这个包下

import (
	"fmt"
)

var num int = initGlobalVariables(100) //第一步,初始化全局变量

func initGlobalVariables(n int) int {

	fmt.Println("初始化全局变量!")

	return n
}

func init() { //第二步调用init()函数

	fmt.Println("执行init函数!")
}

func main() { //第三步调用main函数

}

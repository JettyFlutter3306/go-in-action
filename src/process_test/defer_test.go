package process_test

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(add(30, 60))
}

func add(num1 int, num2 int) int {

	/*
		在Golang中,程序遇到defer关键字,不会立即执行defer后的语句
		而是要将defer后的语句压入一个栈中,然后继续执行函数后面的语句
		遇到defer关键字,会将后面的代码压入栈中,也会将相关的值拷贝到栈中,不会随着函数后面的变化而变化
		defer应用的场景: 比如你想关闭某个使用的资源,在使用的时候随手defer,因为defer有延迟执行的机制
		类比 HTML 标签 <script defer = "true"></script>  延迟加载
	*/
	defer fmt.Println("num1 =", num1)
	defer fmt.Println("num2 =", num2)

	//栈的特点是先进后出,在函数执行完毕之后,从栈中取出语句开始执行,按照先进后出的规则执行语句
	var sum int = num1 + num2

	fmt.Println("sum =", sum)

	return sum
}

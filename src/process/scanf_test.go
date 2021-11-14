package process

import (
	"fmt"
	"testing"
)

func TestScanF(t *testing.T) {
	/*
		实现功能: 键盘录入学生的年龄,姓名,成绩,是否是VIP
		方式1: Scanln()
		方式2: Scanf()
	*/
	var age int
	//fmt.Println("请输入学生的年龄:")
	//fmt.Scanln(&age)  //传入age的地址的目的,在Scanln()函数中,对地址中的值进行改变的时候,实际外面的age被影响了
	//
	var name string
	//fmt.Println("请输入学生的姓名:")
	//fmt.Scanln(&name)
	//
	var score float32
	//fmt.Println("请输入学生的成绩:")
	//fmt.Scanln(&score)
	//
	var isVIP bool
	//fmt.Println("请输入学生是否是VIP:")
	//fmt.Scanln(&isVIP)
	//
	///*
	//打印上述数据
	// */
	//fmt.Printf("学生的年龄是: %v,姓名为: %v, 成绩是: %v, 是否是VIP: %v", age, name, score, isVIP)

	fmt.Println("请录入学生的年龄,姓名,成绩,是否是VIP,使用空格分割")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Printf("学生的年龄是: %v,姓名为: %v, 成绩是: %v, 是否是VIP: %v", age, name, score, isVIP)
}

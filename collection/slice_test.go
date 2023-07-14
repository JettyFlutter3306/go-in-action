package colllection

import (
	"fmt"
	"testing"
)

/*
切片定义语法:

	var slice = collection[a:b]
*/
func TestSlice(t *testing.T) {
	// 定义数组
	var array = [6]int{3, 6, 9, 10, 15, 16}

	// 切片建立在数组之上
	// 定义一个切片 切片长度是动态变化的所以不用写长度
	// [1:3]切片指的是切出的片段 索引从1开始 到3结束 范围前闭后开[1,3)
	var slice = array[1:3]

	// 输出数组
	fmt.Println("collection: ", array)
	fmt.Println("slice: ", slice)

	/*
		collection:  [3 6 9 10 15 16]
		slice:  [6 9]
	*/
	fmt.Println("切片的元素个数: ", len(slice))
	fmt.Println("切片的容量为: ", cap(slice))
}

func TestMakeSlice(t *testing.T) {
	// 使用内置函数make()创建切片
	// make()的3个参数,切片类型,切片长度,切片的容量
	// make()底层创建了一个数组,外界不可访问,只能通过slice间接地操作数组
	slice := make([]int, 4, 20)

	fmt.Println(slice) // [0 0 0 0]

	// 第三种方式与make()类似,但是底层的数组对外界不可见
	slice1 := []int{1, 4, 7}
	fmt.Println(slice1)
}

// 测试切片的遍历
func TestTraverseSlice(t *testing.T) {
	slice := []int{2, 8, 89, 45}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]: %v \t", i, slice[i])
	}

	fmt.Println("-------------------------")

	for i, value := range slice {
		fmt.Printf("%v: %v \t", i, value)
	}

	fmt.Println()
}

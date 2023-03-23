package colllection

import (
	"fmt"
	"testing"
)

/**
数组
*/

// 测试声明数组
func TestStateArray(t *testing.T) {
	// 第一种
	var array1 [3]int = [3]int{3, 5, 9}
	fmt.Println(array1)

	// 第二种
	var array2 = [3]int{3, 9, 45}
	fmt.Println(array2)

	// 第三种
	var array3 = [...]int{4, 5, 6, 7}
	fmt.Println(array3)

	// 第四种
	var array4 = [...]int{2: 66, 0: 33, 1: 99, 3: 88}
	fmt.Println(array4)
}

// 测试一维数组
func TestArray(t *testing.T) {
	var scores [5]int
	var count = 0

	for i := 0; i < len(scores); i += 1 {
		scores[i] = i*3 + 6
		count += scores[i]
	}

	// 使用for-range方式遍历数组,想要忽略某个参数只需要用_代替即可
	for i, score := range scores {
		fmt.Printf("第 %v 个学生成绩为: %v \n", i+1, score)
	}

	fmt.Printf("总和: %v \n", count)
	fmt.Printf("平均数: %v \n", count/len(scores))
	fmt.Printf("数组的类型: %T \n", scores) // [5]int[3 5 9]
}

// 测试二维数组遍历
func Test2DArray(t *testing.T) {
	var arr = [3][3]int{
		{1, 4, 6},
		{2, 5, 8},
		{3, 6, 9},
	}

	fmt.Println(arr)
	fmt.Println("--------------------")

	// 方式1.普通for循环
	for i := 0; i < len(arr); i += 1 {
		for j := 0; j < len(arr[0]); j += 1 {
			fmt.Print(arr[i][j], "\t")
		}

		fmt.Println()
	}

	fmt.Println("------------------------")
	// for - range循环
	for i, values := range arr {
		for j, value := range values {
			fmt.Printf("arr[%v][%v]: %v \t", i, j, value)
		}

		fmt.Println()
	}
}

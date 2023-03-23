package generic_type

import "fmt"

// 定义泛型类型函数
func DefGeneric() {
	// 泛型切片
	type ISlice[T int | string] []T

	// 泛型映射
	type IMap[K int | string, V float32 | float64] map[K]V

	// 泛型结构体
	type IStruct[T int | float64] struct {
		data          []T
		l             int
		max, min, avg T
	}

	// 声明泛型类型变量
	fmt.Println("打印泛型切片:")
	slice1 := ISlice[int]{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	slice2 := ISlice[string]{"1", "2", "3", "4"}
	fmt.Println(slice2)

	fmt.Println("打印泛型映射")
	hashMap := IMap[int, float32]{1: 5.0, 2: 6.21}
	fmt.Println(hashMap)

	fmt.Println("打印泛型结构体")
	s1 := IStruct[float64]{}
	fmt.Println(s1)

	// 打印数据类型
	fmt.Printf("%T\n", slice1)
	fmt.Printf("%T\n", slice2)
	fmt.Printf("%T\n", hashMap)
	fmt.Printf("%T\n", s1)
}

func ExtGeneric() {
	// 类型为基础类型
	type intSlice[T ~int] []T

	_ = intSlice[int]{}

	type myInt int
	type yourInt int

	_ = intSlice[myInt]{}
	_ = intSlice[yourInt]{}

	// 必须要用以int为基础的类型
	// _ = intSlice[int32]{}

}

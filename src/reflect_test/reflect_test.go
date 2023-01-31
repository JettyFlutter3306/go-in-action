package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 利用一个函数,函数的参数定义为空接口
func reflectInt(i interface{}) {
	refType := reflect.TypeOf(i)
	fmt.Println(refType)

	refValue := reflect.ValueOf(i)
	fmt.Println(refValue)
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Println("名字: ", s.Name)
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func reflectStruct(i interface{}) {
	refType := reflect.TypeOf(i)
	fmt.Println(refType)
	fmt.Printf("具体类型: %T\n", refType)

	refValue := reflect.ValueOf(i)
	fmt.Println(refValue)
	fmt.Printf("具体类型: %T\n", refValue)

	// 获取类别
	kind1 := refType.Kind()
	kind2 := refValue.Kind()
	fmt.Printf("kind1: %v, kind2: %v \n", kind1, kind2)

	// 得到结构体内部字段
	n := refValue.NumField()
	fmt.Println("字段数: ", n)
	for i := 0; i < n; i++ {
		fmt.Printf("第 %d 个字段的值为: %v\n", i, refValue.Field(i))
	}

	n1 := refValue.NumMethod()
	fmt.Println("方法数: ", n1)

	// 反射调用方法,方法名称首字母必须要大写
	refValue.MethodByName("Print").Call(nil)

	// 转成接口
	i2 := refValue.Interface()

	// 类型断言
	student, flag := i2.(Student)
	if flag {
		fmt.Printf("Name: %v, Age: %v \n", student.Name, student.Age)
	}
}

func TestReflectBasicType(t *testing.T) {
	var num = 100
	reflectInt(num)
}

func TestReflectStruct(t *testing.T) {
	var student = Student{
		Name: "洛必达",
		Age:  100,
	}
	reflectStruct(student)
}

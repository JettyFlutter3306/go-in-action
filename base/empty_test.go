package base

import (
	"fmt"
	"reflect"
	"testing"
)

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	reValue := reflect.ValueOf(i)

	k1 := reType.Kind()
	fmt.Println(k1)
	k2 := reValue.Kind()
	fmt.Println(k2)

	i2 := reValue.Interface()
	n, flag := i2.(Student)
	if flag {
		fmt.Printf("结构体的类型是：%T", n)
	}
}

type Student struct {
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	stu := Student{
		Name: "丽丽",
		Age:  18,
	}
	testReflect(stu)
}

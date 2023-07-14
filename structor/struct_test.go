package structor_test

import (
	"fmt"
	"testing"
)

// 声明结构体,属性名称和函数名称一样,开头大写表示对外公开,否则私有
type Teacher struct {
	Name   string
	Age    int
	School string
}

func TestStruct(t *testing.T) {
	// 创建老师结构体的实例,变量
	var maShiBing = Teacher{"马士兵", 20, "清华大学"}
	fmt.Println(maShiBing)

	// 声明一个结构体指针,使用指针指向对象的属性赋值
	var p = new(Teacher)
	(*p).Name = "马士兵" // 简化写法 p.Name = "马士兵", 编译器会自动识别
	(*p).Age = 50
	(*p).School = "北京大学"

	// 打印解引用值
	fmt.Println(*p)

	// 最后一种方式
	var p1 = &Teacher{"马士兵", 48, "清华大学"}
	fmt.Println(p1)
}

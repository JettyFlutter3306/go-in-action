package structor

import (
	"fmt"
	"testing"
)

// 定义Person结构体
type Person struct {
	Name string
}

// 给Person结构体绑定方法test 这里是值传递
func (p Person) deliverValue() {
	p.Name = "牛顿"
	fmt.Println(p.Name)
}

// 这里是引用传递
func (p *Person) deliverReference() {
	p.Name = "牛顿"
	fmt.Println(p.Name)
}

func TestPeron(t *testing.T) {
	var p Person
	p.Name = "洛必达"
	fmt.Println(p.Name)
	p.deliverValue()

	p.deliverReference()
	fmt.Println(p.Name)
}

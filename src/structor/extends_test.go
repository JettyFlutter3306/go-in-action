package structor

import (
	"fmt"
	"testing"
)

type Animal struct {
	Age    int
	Weight float32
}

func (a *Animal) Shout() {
	fmt.Println("I can shout")
}

func (a *Animal) ShowInfo() {
	fmt.Printf("动物年龄: %v, 体重是: %v\n", a.Age, a.Weight)
}

// 为了复用性,体现继承思想,加入匿名结构体
type Cat struct {
	Animal
}

// 绑定Cat的特有方法
func (Cat) Scratch() {
	fmt.Println("I am a cat, I can scratch")
}

func TestObject(t *testing.T) {
	cat := &Cat{}
	cat.Age = 3
	cat.Weight = 10.5
	cat.Shout()
	cat.ShowInfo()
	cat.Scratch()
}

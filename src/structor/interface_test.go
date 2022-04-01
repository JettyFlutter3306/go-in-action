package structor

import (
	"fmt"
	"testing"
)

// 定义一个接口
type Hello interface {
	sayHello()
}

type Chinese struct {
}

func (Chinese) sayHello() {
	fmt.Println("你好")
}

type American struct {
}

func (American) sayHello() {
	fmt.Println("Hello Motherfucker")
}

func Greet(h Hello) {
	h.sayHello()
}

func TestInterface(t *testing.T) {
	chinese := Chinese{}
	american := American{}

	Greet(chinese)
	Greet(american)
}

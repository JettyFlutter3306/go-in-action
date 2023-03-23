package structor_test

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

func (Chinese) eatDumplings() {
	fmt.Println("中国人吃饺子,美国人吃骨头")
}

type American struct {
}

func (American) sayHello() {
	fmt.Println("Hello Motherfucker")
}

func Greet(h Hello) {
	h.sayHello()

	// 使用断言判断类型,看看是否能转成Chinese类型
	if ch, ok := h.(Chinese); ok {
		ch.eatDumplings()
	}
}

func TestInterface(t *testing.T) {
	chinese := Chinese{}
	american := American{}

	Greet(chinese)
	Greet(american)
}

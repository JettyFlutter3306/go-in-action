package error

import (
	"errors"
	"fmt"
	"testing"
)

/*
	通过使用 defer + recover() 进行错误处理

*/
func TestRecover(t *testing.T) {
	//利用defer和recover()捕获异常
	defer func() {
		//调用recover()内置函数,可以捕获异常
		err := recover()

		//如果没有捕获异常返回值为零值 nil
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是: ", err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2

	fmt.Println(result)
}

func TestErr(t *testing.T) {
	err := testErr()

	if err != nil {
		fmt.Println("自定义错误: ", err)
	}

	fmt.Println("上面的除法操作执行成功...")
	fmt.Println("正常执行下面的逻辑..")
}

/**
自定义错误
*/
func testErr() (err error) {
	num1 := 10
	num2 := 0

	if num2 == 0 {
		//抛出自定义错误
		return errors.New("除数不可以是0~~")
	} else { //如果除数不为0,那么正常进行就可以了
		result := num1 / num2
		fmt.Println(result)

		//如果没有错误那么就返回nil
		return nil
	}
}

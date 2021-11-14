package process

import (
	"fmt"
	"testing"
)

//Golang中不用加break跳出switch 执行完毕一个分支之后会自动跳出不会出现继续向下执行的情况
func TestSwitch(t *testing.T) {
	var score = 87

	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
	case 8:
		fmt.Println("C")
	case 7:
		fmt.Println("D")
	case 6:
		fmt.Println("E")
	default:
		fmt.Println("F")
	}
}

package string

import (
	"fmt"
	"testing"
)

func TestConvertRadix(t *testing.T) {
	/*
		基本数据类型转为string类型
	*/
	var n1 = 19
	var n2 float32 = 4.78
	var n3 = false
	var n4 byte = 'a'

	s1 := fmt.Sprintf("%d", n1)
	fmt.Printf("类型: %T, s1 = %q \n", s1, s1)

	s2 := fmt.Sprintf("%f", n2)
	fmt.Printf("类型: %T, s2 = %q \n", s2, s2)

	s3 := fmt.Sprintf("%t", n3)
	fmt.Printf("类型: %T, s3 = %q \n", s3, s3)

	s4 := fmt.Sprintf("%c", n4)
	fmt.Printf("类型: %T, s4 = %q \n", s4, s4)
}

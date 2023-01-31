package test_leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	fmt.Println(isValid(s))
	s = "()"
	fmt.Println(isValid(s))
	s = "{]}[]"
	fmt.Println(isValid(s))
	s = ")"
	fmt.Println(isValid(s))
}

package test_leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	x := 123
	y := 121
	fmt.Println(isPalindrome(x))
	fmt.Println(isPalindrome(y))
}

func TestReverse(t *testing.T) {
	x := 123
	fmt.Println(reverse(x))
	x = -123
	fmt.Println(reverse(x))
	x = 120
	fmt.Println(reverse(x))
}

package base

import (
	"fmt"
	"testing"
)

func Test_sortArray(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(sortArray(nums))
}

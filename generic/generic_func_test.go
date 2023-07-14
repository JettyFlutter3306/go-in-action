package generic

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	// 支持类型推断
	fmt.Println(Sum(1, 2, 4, 5, 10))
	fmt.Println(Sum[string]("You", "FOOL", "Fucking", "Stupid"))
}

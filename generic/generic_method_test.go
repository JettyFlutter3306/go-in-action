package generic

import (
	"fmt"
	"testing"
)

func TestGenericList(t *testing.T) {
	l := List[int]{}
	l.Add(1).Add(2).Add(10).Add(20)
	fmt.Println(l.Data())
	fmt.Println("max: ", l.Max())
	fmt.Println("min: ", l.Min())
}

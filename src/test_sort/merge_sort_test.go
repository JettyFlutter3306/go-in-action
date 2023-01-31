package test_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	MergeSort(arr)
	fmt.Println(arr)
}

func TestMerge(t *testing.T) {
	l := 0
	r := len(arr) - 1
	m := l + ((r - l) >> 1)
	merge(arr, l, m, r)
	fmt.Println(arr)
}

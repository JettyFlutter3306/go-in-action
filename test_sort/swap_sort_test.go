package test_sort

import (
	"fmt"
	"testing"
)

var arr = []int{2, 10, 7, 28, 7, 19, 10, 8, 4, 7, 7}

func TestBubbleSort(t *testing.T) {
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	QuickSort(arr)
	fmt.Println(arr)
}

func TestPartition(t *testing.T) {
	index := partition(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println(index)
}

func TestNetherlandsFlag(t *testing.T) {
	section := netherlandsFlag(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println(section)
}

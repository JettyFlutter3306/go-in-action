package sort

import (
	"fmt"
	"testing"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	splitData := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	high := make([]int, 0, 0)    //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitData) //加入一个

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, high = QuickSort(low), QuickSort(high)
	ans := append(append(low, mid...), high...)

	return ans
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}

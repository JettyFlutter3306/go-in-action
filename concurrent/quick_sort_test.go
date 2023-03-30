package concurrent

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var arr []int
	max := 1_000_000_00
	rand.Seed(time.Now().Unix())
	for i := 0; i < max; i++ {
		arr = append(arr, rand.Intn(max))
	}
	QuickSort(arr)
}

func TestQuickSortConcurrent(t *testing.T) {
	var arr []int
	max := 1_000_000_00
	rand.Seed(time.Now().Unix())
	for i := 0; i < max; i++ {
		arr = append(arr, rand.Intn(max))
	}
	QuickSortConcurrent(arr)
}

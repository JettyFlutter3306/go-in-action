package concurrent

import "sync"

func QuickSortConcurrent(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go quickSortConcurrent(arr, 0, len(arr)-1, wg)
	wg.Wait()
}

func quickSortConcurrent(arr []int, l, r int, wg *sync.WaitGroup) {
	defer wg.Done()
	if l < r {
		mid := partition(arr, l, r)
		wg.Add(2)
		go quickSortConcurrent(arr, l, mid-1, wg)
		go quickSortConcurrent(arr, mid+1, r, wg)
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		mid := partition(arr, l, r)
		quickSort(arr, l, mid-1)
		quickSort(arr, mid+1, r)
	}
}

// 将数组划分为三块，小于基准值、等于基准值、大于基准值
func partition(arr []int, l, r int) int {
	less := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= arr[r] {
			less++
			arr[i], arr[less] = arr[less], arr[i]
		}
	}
	return less
}

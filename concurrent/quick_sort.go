package concurrent

import "sync"

func QuickSortConcurrent(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go quickSortConcurrent(arr, 0, len(arr)-1, wg, 1)
	wg.Wait()
}

const maxDeep = 9

func quickSortConcurrent(arr []int, l, r int, wg *sync.WaitGroup, dp int) {
	if dp <= maxDeep+1 {
		defer wg.Done()
	}

	if l < r {
		mid := partition(arr, l, r)
		if dp <= maxDeep {
			wg.Add(2)
			go quickSortConcurrent(arr, l, mid-1, wg, dp+1)
			go quickSortConcurrent(arr, mid+1, r, wg, dp+1)
		} else {
			quickSortConcurrent(arr, l, mid-1, wg, dp+1)
			quickSortConcurrent(arr, mid+1, r, wg, dp+1)
		}
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

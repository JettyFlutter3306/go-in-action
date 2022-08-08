package test_sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	var exchanged = true // 是否需要交换的标志
	for i := 0; i < len(arr)-1 && exchanged; i++ {
		exchanged = false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				exchanged = true
			}
		}
	}
}

// 快速排序
func QuickSort(arr []int) {
	ProcessQuickSort(arr, 0, len(arr)-1)
}

func ProcessQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	section := netherlandsFlag(arr, l, r)
	ProcessQuickSort(arr, l, section[0]-1)
	ProcessQuickSort(arr, section[1]+1, r)
}

// 快速排序的划分辅助函数 假如存在多个与基准值相等的元素需要进行重复计算
// 返回基准值的索引
func partition(arr []int, l, r int) int {
	if l > r {
		return -1
	} else if l == r {
		return 1
	}

	lessEqual := l - 1
	index := l

	// 声明index永远指向小于基准值区间的末尾
	for index < r {
		if arr[index] <= arr[r] {
			lessEqual++
			arr[index], arr[lessEqual] = arr[lessEqual], arr[index]
		}
		index++
	}

	// 交换基准值和大于基准值的第一个数
	lessEqual++
	arr[lessEqual], arr[r] = arr[r], arr[lessEqual]

	return index
}

// 将数组分成三个部分返回中间基准值的区间数组，也就是荷兰国旗问题
func netherlandsFlag(arr []int, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	} else if l == r {
		return []int{l, r}
	}

	less := l - 1 // 小于基准数的右边界
	more := r     // 大于基准指数的左边界
	index := l

	for index < more {
		if arr[index] == arr[r] {
			index++
		} else if arr[index] < arr[r] {
			arr[less+1], arr[index] = arr[index], arr[less+1]
			less++
			index++
		} else {
			more--
			arr[index], arr[more] = arr[more], arr[index]
		}
	}

	arr[more], arr[r] = arr[r], arr[more]
	return []int{less + 1, more}
}

package test_sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	m := l + ((r - l) >> 1)

	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

// 归并排序辅助函数
func merge(arr []int, l, m, r int) {
	helper := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := m + 1

	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
			helper[i] = arr[p1]
			p1++
		} else {
			helper[i] = arr[p2]
			p2++
		}

		i++
	}

	// 要么p1越界了,要么p2越界了,开始拷贝剩下的数据到helper数组里面,两个for只会发生一个
	for p1 <= m {
		helper[i] = arr[p1]
		i++
		p1++
	}

	for p2 <= r {
		helper[i] = arr[p2]
		i++
		p2++
	}

	// 把helper数组的元素拷贝回原数组中去
	for index := range helper {
		arr[l+index] = helper[index]
	}
}

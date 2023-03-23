package test_leetcode

/*
NO.35  简单  搜索插入的位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
其实也就是求出大于等于target的最小索引!
输入: nums = [1,3,5,6], target = 5
输出: 2
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	ans := n

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

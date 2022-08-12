package test_leetcode

/*
NO.1  简单  两数之和
给定一个整数数组 nums和一个整数目标值 target，
请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 使用哈希表优化时间复杂度
func twoSum1(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		if p, ok := hashMap[target-v]; ok {
			return []int{p, i}
		}
		hashMap[v] = i
	}

	return nil
}

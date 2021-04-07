/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return 
	}

	mid := nums[(start + end) / 2]  // 随便选取一个值
	l, r := start, end

	for l <= r { // 为了后面遍历的时候 r在l前 分开两个分区 这里带=
		for nums[l] < mid {
			l++
		}
		for nums[r] > mid {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	quickSort(nums, start, r)
	quickSort(nums, l, end)
}
// @lc code=end


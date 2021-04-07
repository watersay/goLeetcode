/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
var res int
func findKthLargest(nums []int, k int) int {
	// 小顶堆 保留k个数 最大的里面k个数 第k大的就再堆顶

	// 快速选择法
	res = -1
	target := len(nums) - k // 把第k大的数 转换为索引位置

	helper(nums, target, 0, len(nums) - 1)
	return res
}

func helper(nums []int, target, start, end int) {
	if start > end { // == 的情况 需要往下走 判断是否是target坐标
		return
	}

	val := nums[start] // 选取第一个
	l, r := start + 1, end // 因为选取第一个 所以不让过start

	for l <= r {
		for l <= r && nums[l] < val { // 前面加上条件 防止溢出
			l++
		}
		for l <= r && nums[r] > val {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	nums[start], nums[r] = nums[r], nums[start] // 把val的值放到正确的位置上

	if r == target {
		res = nums[r]
		return
	} else if r < target {
		helper(nums, target, r + 1, end)
	} else {
		helper(nums, target, start, r - 1)
	}
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums) - 1
	tmp := -1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			tmp = mid
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	if tmp == -1 {
		return []int{-1, -1}
	}

	rl, rr := tmp, tmp
	for rl - 1 >= 0 {
		if nums[rl - 1] == target {
			rl--
		} else {
			break
		}
	}
	for rr + 1 <= len(nums) - 1 {
		if nums[rr + 1] == target {
			rr++
		} else {
			break
		}
	}

	return []int{rl, rr}
}
// @lc code=end


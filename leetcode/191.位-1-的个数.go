/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	
	for num != 0 {
		num = num & (num - 1) // 去掉二进制的最后一个1
		res++
	}

	return res
}
// @lc code=end


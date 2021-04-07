/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// 栈判断

	stack := []rune{} // for 遍历的v 是rune类型！
	dist := map[rune]rune{
		')' : '(',
		']' : '[',
		'}' : '{',
	}

	for _, v := range s {
		if r, ok := dist[v]; ok {
			// 闭合符号
			if len(stack) == 0 { // 栈内空 直接返回
				return false
			}

			pop := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			if pop != r {
				return false
			}
		} else {
			// 开始符号
			stack = append(stack, v)
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true

}
// @lc code=end


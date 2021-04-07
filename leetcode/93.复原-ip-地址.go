/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
var res []string
func restoreIpAddresses(s string) []string {
	// 回溯算法
	res := make([]string, 0)
}

func helper(s string, index int, ip string, part int){
	if part == 4 && i == len(s){ // 符合规则的
		res = append(res, ip)
		return
	}

	if (part == 4 && i < len(s)) || i >= len(s){ //多了字符
		return
	}

	for 

}
// @lc code=end


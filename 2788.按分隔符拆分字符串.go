/*
 * @lc app=leetcode.cn id=2788 lang=golang
 *
 * [2788] 按分隔符拆分字符串
 */
package leetcode

// @lc code=start
func splitWordsBySeparator(words []string, separator byte) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		index := 0
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] == separator {
				if index == j {
					index++
				} else {
					res = append(res, string(words[i][index:j]))
					index = j + 1
				}
			} else if j == len(words[i])-1 && index != j+1 {
				res = append(res, string(words[i][index:j+1]))
			}
		}
	}
	return res
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	slen := len(s)
	tlen := len(t)
	if slen < tlen {
		return 0
	}
	res := make([][]int, slen+1)
	for i := range res {
		res[i] = make([]int, tlen+1)
	}
	for i := 0; i < slen+1; i++ {
		res[i][tlen] = 1
	}
	for i := slen - 1; i >= 0; i-- {
		for j := tlen - 1; j >= 0; j-- {
			if s[i] == t[j] {
				res[i][j] = res[i+1][j+1] + res[i+1][j]
			} else {
				res[i][j] = res[i+1][j]
			}
		}
	}
	return res[0][0]
}

// @lc code=end


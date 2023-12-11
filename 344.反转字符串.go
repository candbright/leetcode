/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		tmp := s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
}

// @lc code=end


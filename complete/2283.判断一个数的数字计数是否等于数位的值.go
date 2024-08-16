/*
 * @lc app=leetcode.cn id=2283 lang=golang
 *
 * [2283] 判断一个数的数字计数是否等于数位的值
 */
package leetcode

// @lc code=start
func digitCount(num string) bool {
	c := make(map[int]int)
	for _, r := range num {
		c[int(r-48)]++
	}
	for i := 0; i < len(num); i++ {
		should := int(num[i] - 48)
		if should != c[i] {
			return false
		}
	}
	return true
}

// @lc code=end

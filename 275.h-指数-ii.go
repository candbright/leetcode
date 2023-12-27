/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */
package leetcode

// @lc code=start
func hIndex(citations []int) int {
	for i := 0; i < len(citations); i++ {
		index := len(citations) - 1 - i
		if citations[index] < i+1 {
			return i
		}
		if index == 0 {
			return i + 1
		}

	}
	return 0
}

// @lc code=end

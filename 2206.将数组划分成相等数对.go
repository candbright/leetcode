/*
 * @lc app=leetcode.cn id=2206 lang=golang
 *
 * [2206] 将数组划分成相等数对
 */

// @lc code=start
func divideArray(nums []int) bool {
	div := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := div[nums[i]]
		if !ok {
			div[nums[i]] = 1
		} else {
			delete(div, nums[i])
		}
	}
	return len(div) == 0
}

// @lc code=end


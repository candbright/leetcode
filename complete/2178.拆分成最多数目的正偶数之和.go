/*
 * @lc app=leetcode.cn id=2178 lang=golang
 *
 * [2178] 拆分成最多数目的正偶数之和
 */

// @lc code=start
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return []int64{}
	}
	result := make([]int64, 0)
	cur := int64(2)
	sum := int64(0)
	for {
		sum += cur
		if finalSum >= sum+cur+2 {
			result = append(result, cur)
		} else {
			sum -= cur
			result = append(result, finalSum-sum)
			break
		}
		cur += 2
	}
	return result
}

// @lc code=end


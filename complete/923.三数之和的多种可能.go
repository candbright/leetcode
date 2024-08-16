/*
 * @lc app=leetcode.cn id=923 lang=golang
 *
 * [923] 三数之和的多种可能
 */
package leetcode

// @lc code=start
func threeSumMulti(arr []int, target int) int {
	count := make(map[int]int)
	res := 0
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for k1 := range count {
		count1 := count[k1]
		count[k1]--
		for k2 := range count {
			count2 := count[k2]
			if count2 == 0 {
				continue
			}
			count[k2]--
			third := target - k1 - k2
			if count[third] > 0 {
				res += count[third] * count2 * count1
			}
			count[k2]++
		}
		count[k1]++
	}
	return res / 6 % (1e9 + 7)
}

// @lc code=end

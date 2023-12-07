/*
 * @lc app=leetcode.cn id=1647 lang=golang
 *
 * [1647] 字符频次唯一的最小删除次数
 */
package leetcode

// @lc code=start
func minDeletions(s string) int {
	num := make([]int, 26)
	for _, c := range s {
		num[c-97]++
	}
	res := 0
	set := make(map[int]struct{})
	for i := 0; i < len(num); i++ {
		if num[i] == 0 {
			continue
		}
		val := num[i]
		_, ok := set[val]
		if !ok {
			set[val] = struct{}{}
			continue
		} else {
			for {
				val--
				res++
				if val == 0 {
					break
				}
				_, ok := set[val]
				if !ok {
					set[val] = struct{}{}
					break
				}
			}
		}
	}
	return res
}

// @lc code=end

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
	unique := make([]int, 0)
	for i := 0; i < len(num); i++ {
		if num[i] == 0 {
			continue
		}
		index := 0
		for {
			if index == len(unique) {
				unique = append(unique, num[i])
				break
			}
			if num[i] > unique[index] {
				left := make([]int, len(unique[:index]))
				copy(left, unique[:index])
				right := make([]int, len(unique[index:]))
				copy(right, unique[index:])
				unique = append(left, num[i])
				unique = append(unique, right...)
				break
			} else if num[i] == unique[index] {
				if num[i] == 0 {
					break
				}
				num[i]--
				index++
				res++
			} else {
				index++
			}
		}
	}
	return res
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=1647 lang=golang
 *
 * [1647] 字符频次唯一的最小删除次数
 */
package leetcode

// @lc code=start
func minDeletions(s string) int {
	//0-25的值分别表示ascii码97-122出现的次数
	num := make([]int, 26)
	for _, c := range s {
		num[c-97]++
	}
	res := 0
	////key表示已经插入的次数，value无意义
	set := make(map[int]struct{})
	for i := 0; i < len(num); i++ {
		if num[i] == 0 {
			continue
		}
		for {
			//次数减为0时跳出循环
			if num[i] == 0 {
				break
			}
			_, ok := set[num[i]]
			////当前字符的次数没被占用，直接插入
			if !ok {
				set[num[i]] = struct{}{}
				break
			}
			//当前字符的次数被占用，次数减一，统计结果加一，并重新判断
			num[i]--
			res++
		}
	}
	return res
}

// @lc code=end

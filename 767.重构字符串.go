/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */
package leetcode

// @lc code=start
func reorganizeString(s string) string {
	var max, maxIndex int
	count := make([]int, 26)
	for _, c := range s {
		count[c-97]++
		if count[c-97] > max {
			max = count[c-97]
			maxIndex = int(c - 97)
		}
	}
	remain := 0
	for i := 0; i < 26; i++ {
		if i != maxIndex {
			remain += count[i]
		}
	}
	if max-1 > remain {
		return ""
	} else {
		res := []rune(s)
		i := 0
		for {
			if count[maxIndex] == 0 {
				break
			}
			res[i] = rune(maxIndex + 97)
			count[maxIndex]--
			i += 2
		}
		j := 0
		for {
			if i >= len(res) {
				break
			}
			for {
				if count[j] == 0 {
					j++
				} else {
					break
				}
			}
			res[i] = rune(j + 97)
			count[j]--
			i += 2
		}
		i = 1
		for {
			if i >= len(res) {
				break
			}
			for {
				if count[j] == 0 {
					j++
				} else {
					break
				}
			}
			res[i] = rune(j + 97)
			count[j]--
			i += 2
		}
		return string(res)
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */
package leetcode

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	temp := pushed[:1]
	pushedIndex := 1
	poppedIndex := 0
	for {
		if len(temp) == 0 {
			if pushedIndex <= len(pushed)-1 {
				temp = append(temp, pushed[pushedIndex])
				pushedIndex++
			}
		} else {
			if temp[len(temp)-1] != popped[poppedIndex] {
				if pushedIndex <= len(pushed)-1 {
					temp = append(temp, pushed[pushedIndex])
					pushedIndex++
				} else {
					return false
				}
			} else {
				temp = temp[:len(temp)-1]
				if len(temp) == 0 && poppedIndex == len(popped)-1 {
					return true
				} else {
					poppedIndex++
				}
			}
		}

	}
}

// @lc code=end

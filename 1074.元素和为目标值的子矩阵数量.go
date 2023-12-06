/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */
package leetcode

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	lenX := len(matrix)
	lenY := len(matrix[0])
	res := 0
	m := make([][]int, lenX)
	for i := 0; i < lenX; i++ {
		m[i] = make([]int, lenY)
	}
	for startX := 0; startX < lenX; startX++ {
		for startY := 0; startY < lenY; startY++ {
			for curX := startX; curX < lenX; curX++ {
				sum := 0
				for curY := startY; curY < lenY; curY++ {
					sum += matrix[curX][curY]
					if curX == startX {
						m[curX][curY] = sum
					} else {
						m[curX][curY] = sum + m[curX-1][curY]
					}
					if m[curX][curY] == target {
						res++
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

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
	for startY := 0; startY < lenY; startY++ {
		x := 0
		y := startY
		for x < lenX && y < lenY {
			if x == lenX {
				x = 0
				y++
			}
			for curX := x; curX < lenX; curX++ {
				sum := 0
				for curY := y; curY < lenY; curY++ {
					sum += matrix[curX][curY]
					if curX == x {
						m[curX][curY] = sum
					} else {
						m[curX][curY] = sum + m[curX-1][curY]
					}
					if m[curX][curY] == target {
						res++
					}
				}
			}
			x++
		}
	}
	return res
}

// @lc code=end

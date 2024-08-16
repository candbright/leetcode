/*
 * @lc app=leetcode.cn id=1266 lang=golang
 *
 * [1266] 访问所有点的最小时间
 */
package leetcode

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		res += calculateTime(points[i-1], points[i])
	}
	return res
}

func calculateTime(point1, point2 []int) int {
	x := point2[0] - point1[0]
	y := point2[1] - point1[1]
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end

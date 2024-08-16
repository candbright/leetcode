/*
 * @lc app=leetcode.cn id=2787 lang=golang
 *
 * [2787] 将一个数字表示成幂的和的方案数
 */
package leetcode

// @lc code=start
func numberOfWays(n int, x int) int {
	res := make([]int, n+1)
	res[0] = 1
	i := 1
	for {
		p := pow(i, x)
		if p > n {
			break
		}
		for s := n; s >= p; s-- {
			res[s] += res[s-p]
		}
		i++
	}
	return res[n] % (1e9 + 7)
}

func pow(i, x int) int {
	res := 1
	for ; x > 0; x /= 2 {
		if x%2 > 0 {
			res = res * i
		}
		i = i * i
	}
	return res
}

// @lc code=end

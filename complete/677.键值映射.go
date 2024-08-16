/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */
package leetcode

// @lc code=start
type MapSum struct {
	data map[string]int
}

func Constructor() MapSum {
	return MapSum{
		data: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for key, val := range this.data {
		if contains(key, prefix) {
			res += val
		}
	}
	return res
}

func contains(a, b string) bool {
	if len(a) < len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

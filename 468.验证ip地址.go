/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */
package leetcode

import (
	"strconv"
	"strings"
)

// @lc code=start
func validIPAddress(queryIP string) string {
	split4 := strings.Split(queryIP, ".")
	split6 := strings.Split(queryIP, ":")
	if len(split4) == 4 {
		//ipv4
		for _, s := range split4 {
			i, err := strconv.Atoi(s)
			if err != nil {
				return "Neither"
			}
			if i > 255 {
				return "Neither"
			}
			if s[0] == 48 && (i != 0 || len(s) > 1) {
				return "Neither"
			}
		}
		return "IPv4"
	} else if len(split6) == 8 {
		if len(split6) != 8 {
			return "Neither"
		}
		for _, s := range split6 {
			//0-9	48-57
			//a-f	97-102
			//A-F	65-70
			if len(s) != 0 && len(s) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(s, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	} else {
		return "Neither"
	}
}

// @lc code=end

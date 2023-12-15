/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
package leetcode

// @lc code=start
func multiply(num1 string, num2 string) string {
	num1Len := len(num1)
	if num1Len == 1 && num1[0] == 48 {
		return "0"
	}
	num2Len := len(num2)
	if num2Len == 1 && num2[0] == 48 {
		return "0"
	}
	res := "0"
	for i := num2Len - 1; i >= 0; i-- {
		tmp := make([]int, num1Len+1)
		carry := 0
		for j := num1Len - 1; j >= 0; j-- {
			m := getNum(num1[j]) * getNum(num2[i])
			carry = m / 10
			remain := m - 10*carry
			tmp[j+1] += remain
			if tmp[j+1] > 9 {
				tmp[j+1] -= 10
				tmp[j]++
			}
			tmp[j] += carry
		}
		for k := 0; k < num2Len-1-i; k++ {
			tmp = append(tmp, 0)
		}
		res = add(res, tmp)
	}
	return res
}

func getNum(b byte) int {
	return int(b) - 48
}

func getByte(i int) byte {
	return byte(i + 48)
}

func removePrefix0(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		if b[i] != 48 {
			return b[i:]
		}
	}
	return b[0:1]
}

func add(s string, ints []int) string {
	tmp := []byte(s)
	for i := 0; i < len(ints); i++ {
		sIndex := len(tmp) - 1 - i
		iIndex := len(ints) - 1 - i
		if sIndex < 0 {
			tmp = append([]byte{getByte(0)}, tmp...)
			sIndex = 0
		}
		m := getNum(tmp[sIndex]) + ints[iIndex]
		if m > 9 {
			tmp[sIndex] = getByte(m - 10)
			sIndex--
			for {
				if sIndex < 0 {
					tmp = append([]byte{getByte(1)}, tmp...)
					break
				}
				tmp[sIndex]++
				if tmp[sIndex] > getByte(9) {
					tmp[sIndex] -= 10
					sIndex--
				} else {
					break
				}
			}
		} else {
			tmp[sIndex] = getByte(m)
		}
	}
	return string(removePrefix0(tmp))
}

// @lc code=end

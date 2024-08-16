/*
 * @lc app=leetcode.cn id=1331 lang=golang
 *
 * [1331] 数组序号转换
 */
package leetcode

// @lc code=start
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	indexArr := make([]int, len(arr))
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indexArr[i] = i
	}
	quickSort(arr, indexArr)
	tmp := 1
	res[indexArr[0]] = tmp
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			res[indexArr[i]] = tmp
		} else {
			tmp += 1
			res[indexArr[i]] = tmp
		}
	}
	return res
}

func quickSort(arr, indexArr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			indexArr[left], indexArr[right] = indexArr[right], indexArr[left]
			left++
			right--
		}
	}

	if right > 0 {
		quickSort(arr[:right+1], indexArr[:right+1])
	}
	if left < len(arr)-1 {
		quickSort(arr[left:], indexArr[left:])
	}
}

// @lc code=end

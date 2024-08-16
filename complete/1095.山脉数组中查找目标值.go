/*
 * @lc app=leetcode.cn id=1095 lang=golang
 *
 * [1095] 山脉数组中查找目标值
 */

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	start := 0
	end := mountainArr.length() - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		midVal := mountainArr.get(mid)
		afterVal := mountainArr.get(mid + 1)
		if start == end {
			break
		}
		if afterVal > midVal {
			start = mid + 1
		}
		if afterVal < midVal {
			end = mid
		}
	}
	val, ok := findInIncreasePart(target, 0, mid, mountainArr)
	if ok {
		return val
	}
	val, ok = findIndDecreasePart(target, mid, mountainArr.length()-1, mountainArr)
	if ok {
		return val
	} else {
		return -1
	}
}

func findInIncreasePart(target, startIndex, endIndex int, mountainArr *MountainArray) (int, bool) {
	midIndex := (startIndex + endIndex) / 2
	midValue := mountainArr.get(midIndex)
	if midValue == target {
		return midIndex, true
	} else {
		if startIndex >= endIndex {
			return -1, false
		} else if midValue > target {
			return findInIncreasePart(target, startIndex, midIndex-1, mountainArr)
		} else {
			return findInIncreasePart(target, midIndex+1, endIndex, mountainArr)
		}
	}
}

func findIndDecreasePart(target, startIndex, endIndex int, mountainArr *MountainArray) (int, bool) {
	midIndex := (startIndex + endIndex) / 2
	midValue := mountainArr.get(midIndex)
	if midValue == target {
		return midIndex, true
	} else {
		if startIndex >= endIndex {
			return -1, false
		} else if midValue < target {
			return findIndDecreasePart(target, startIndex, midIndex-1, mountainArr)
		} else {
			return findIndDecreasePart(target, midIndex+1, endIndex, mountainArr)
		}
	}
}

// @lc code=end


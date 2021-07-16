package learning

// 986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	aIdx, bIdx := 0, 0
	aLen, bLen := len(firstList), len(secondList)
	res := make([][]int, 0, min(aLen, bLen))
	for aIdx < aLen && bIdx < bLen {
		interval, less := getInterval(firstList[aIdx], secondList[bIdx])
		if less {
			aIdx++
		} else {
			bIdx++
		}

		if len(interval) == 2 {
			res = append(res, interval)
		}
	}
	return res
}

// 计算ab两区间交集，并返回a区间的右值是否≤b区间
func getInterval(a, b []int) ([]int, bool) {
	if len(a) != 2 || len(b) != 2 {
		return nil, true
	}

	// 1.不相交的情况
	if a[1] < b[0] {
		return nil, true
	}
	if a[0] > b[1] {
		return nil, false
	}

	// 2.包含的情况
	if a[0] <= b[0] && a[1] >= b[1] {
		return b, false
	}
	if a[0] >= b[0] && a[1] <= b[1] {
		return a, true
	}

	// 3.相交的情况
	if a[1] <= b[1] {
		return []int{b[0], a[1]}, true
	} else {
		return []int{a[0], b[1]}, false
	}
}

// 11. 盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		curLeft, curRight := height[left], height[right]
		cur := (right - left) * min(curLeft, curRight)
		if cur > max {
			max = cur
		}

		if height[left] <= height[right] {
			left++
			for left < right && curLeft >= height[left] {
				left++
			}
		} else {
			right--
			for left < right && curRight >= height[right] {
				right--
			}
		}
	}
	return max
}

// 438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	charCount := make([]int, 26)
	for _, v := range p {
		charCount[v-'a']++
	}

	curCount := make([]int, 26)
	left, right := 0, len(p)-1
	for i := left; i <= right && i < len(s); i++ {
		curCount[s[i]-'a']++
	}
	res := make([]int, 0, len(p))
	if checkEqual(charCount, curCount) {
		res = append(res, 0)
	}

	for right++; right < len(s); right++ {
		curCount[s[left]-'a']--
		curCount[s[right]-'a']++

		left++
		if checkEqual(charCount, curCount) {
			res = append(res, left)
		}
	}
	return res
}
func checkEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// 713. 乘积小于K的子数组 TODO bug
func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	curMultiply, limit := float64(nums[left]), float64(k)
	// 获取最大起始区间 左值
	for left < len(nums) && float64(nums[left]) >= limit {
		left++
	}
	if left >= len(nums) {
		return 0
	}
	right = left
	// 获取最大起始区间 右值
	for right < len(nums)-1 {
		tmp := curMultiply * float64(nums[right+1])
		if tmp >= limit {
			break
		}
		curMultiply = tmp
		right++
	}
	res := getSubarrayNum(right - left + 1)
	// 开始滑动
	for left < len(nums)-1 && right < len(nums) {
		curMultiply /= float64(nums[left])
		left++
		if right < left {
			for left < len(nums) && float64(nums[left]) >= limit {
				left++
			}
			curMultiply = float64(nums[left])
			right = left
		}

		newRight := right
		for newRight < len(nums)-1 {
			tmp := curMultiply * float64(nums[newRight+1])
			if tmp >= limit {
				break
			}
			newRight++
			curMultiply = tmp
		}
		if newRight >= left {
			res += getSubarrayNum(newRight - left + 1)
		}
		if newRight > right {
			res -= getSubarrayNum(getDuplicateInterval(left, right)) // 删除重复区间的重复子数组
			right = newRight
		}
	}
	return res
}
func getDuplicateInterval(newLeft, oldRight int) int {
	if oldRight < newLeft {
		return 0
	}
	return oldRight - newLeft + 1
}
func getSubarrayNum(arrLen int) int {
	return (arrLen * (arrLen + 1)) >> 1
}

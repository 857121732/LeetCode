package learning

// 718.最长重复子数组

// 方法一：动态规划
func findLength1(nums1 []int, nums2 []int) int {
	rows, cols := len(nums1), len(nums2)
	dp := make([][]int, rows)
	for i := rows - 1; i >= 0; i-- {
		dp[i] = make([]int, cols)

		// 赋值最后一列
		if nums1[i] == nums2[cols-1] {
			dp[i][cols-1] = 1
		} else if i < rows-1 {
			dp[i][cols-1] = dp[i+1][cols-1]
		}
	}

	// 赋值最后一行
	for j := cols - 2; j >= 0; j-- {
		if nums1[rows-1] == nums2[j] {
			dp[rows-1][j] = 1
		} else {
			dp[rows-1][j] = dp[rows-1][j+1]
		}
	}

	res := 0
	for i := rows - 2; i >= 0; i-- {
		for j := cols - 2; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}

// 方法二：滑动窗口
func findLength2(nums1 []int, nums2 []int) int {
	res := 0
	for i := range nums1 {
		cur := maxLength(nums1[i:], nums2)
		if cur > res {
			res = cur
		}
	}

	for j := range nums2 {
		cur := maxLength(nums1, nums2[j:])
		if cur > res {
			res = cur
		}
	}

	return res
}

func maxLength(a, b []int) int {
	length := min(len(a), len(b))

	maxLength, curLength := 0, 0
	for i := 0; i < length; i++ {
		if a[i] == b[i] {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 0
		}
	}
	if curLength > maxLength {
		maxLength = curLength
	}
	return maxLength
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// qPow 快速幂取余
func qPow(base, power, mod int) int {
	res := 1
	baseTmp := base
	for power != 0 {
		if power&1 == 1 { // 奇次幂
			res = res * baseTmp % mod
		}
		baseTmp = baseTmp * baseTmp % mod
		power >>= 1
	}
	return res
}

// 方法三：二分 + Robin Karp hash TODO

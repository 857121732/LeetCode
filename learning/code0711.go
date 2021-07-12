package learning

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	for k, v := range nums {
		if v < target {
			continue
		}
		if v > target {
			break
		}

		if res[0] == -1 {
			res[0] = k
		}
		res[1] = k
	}
	return res
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) >> 1
	for left <= right {
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // 左为升，右可能为转
			if nums[left] <= target && nums[mid] > target { // 在左边
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] <= nums[right] { // 右为升，左可能为转
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			break
		}

		mid = (left + right) >> 1
	}

	return -1
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][n-1] < target {
			continue
		}

		return searchArr(matrix[i], target)
	}
	return false
}

func searchArr(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	mid := (left + right) >> 1
	for left <= right {
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) >> 1
	}
	return false
}

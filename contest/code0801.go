package contest

import "sort"

func isThree(n int) bool {
	num := 1
	for i := 2; i <= n; i++ {
		if num > 3 {
			return false
		}
		if n%i == 0 {
			num++
		}
	}
	return num == 3
}

// 5831. 你可以工作的最大周数
func numberOfWeeks(milestones []int) int64 {
	sort.Slice(milestones, func(i, j int) bool {
		return milestones[i] >= milestones[j]
	})

	n := len(milestones)
	if n == 1 {
		return 1
	}

	idxA, idxB, idxPre := 0, 1, -1
	res := 0
	for idxA < n && idxB < n {
		num, idxLast := calc(idxA, idxB, idxPre, milestones)
		idxPre = idxLast
		res += num

		if milestones[idxA] == 0 {
			idxA = max(idxA, idxB) + 1
		}
		if milestones[idxB] == 0 {
			idxB = max(idxA, idxB) + 1
		}
	}
	if (idxA < n && milestones[idxA] > 0 && idxA != idxPre) ||
		(idxB < n && milestones[idxB] > 0 && idxB != idxPre) {
		res++
	}
	return int64(res)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func calc(idxA, idxB, idxPre int, arr []int) (num, idxLast int) {
	// 让A指针作为先被减的那个
	if arr[idxA] < arr[idxB] || idxA == idxPre {
		idxA, idxB = idxB, idxA
	}

	if arr[idxA] > arr[idxB] {
		num = (arr[idxB] << 1) + 1
		arr[idxA] -= arr[idxB] + 1
		arr[idxB] = 0
		idxLast = idxA
	} else {
		num = arr[idxA] << 1
		arr[idxB] -= arr[idxA]
		arr[idxA] = 0
		idxLast = idxB
	}
	return
}

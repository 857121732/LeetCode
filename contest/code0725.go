package contest

import (
	"sort"
)

func getLucky(s string, k int) int {
	num := 0
	for _, v := range s {
		num += int(v) - 'a' + 1
	}

	for k > 1 {
		num = convert(num)
		k--
	}
	return num
}

func convert(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}

func maximumNumber(num string, change []int) string {
	beginChange := false
	res := make([]byte, len(num))
	for k, v := range num {
		res[k] = byte(v)
	}

	for k, v := range res {
		curNum := int(v - '0')
		if beginChange && curNum > change[curNum] {
			break
		}

		if curNum < change[curNum] {
			beginChange = true
			res[k] = byte(change[curNum]) + '0'
		}
	}

	return string(res)
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	score := make([][]int, m)
	stuMaxScore := make(map[int]int, m)
	for stuI := 0; stuI < m; stuI++ {
		score[stuI] = make([]int, m)

		curMax := -1
		for mentorI := 0; mentorI < m; mentorI++ {
			curScore := calcScore(students[stuI], mentors[mentorI])
			score[stuI][mentorI] = curScore
			if curScore > curMax {
				curMax = curScore
			}
		}

		stuMaxScore[stuI] = curMax
	}

	// 根据每行最大值排序
	sort.Slice(score, func(i, j int) bool {
		return stuMaxScore[i] >= stuMaxScore[j]
	})

	mentorFlag := make(map[int]bool, m)
	maxScore := 0
	for stuI := 0; stuI < m; stuI++ {
		curMax, maxIdx := -1, -1
		for mentorI := 0; mentorI < m; mentorI++ {
			if mentorFlag[mentorI] {
				continue
			}
			if maxIdx == -1 || score[stuI][mentorI] > curMax {
				curMax = score[stuI][mentorI]
				maxIdx = mentorI
			}
		}
		mentorFlag[maxIdx] = true
		maxScore += curMax
	}
	return maxScore
}

func calcScore(a, b []int) int {
	if len(a) != len(b) {
		return 0
	}

	num := 0
	for k, v := range a {
		if v == b[k] {
			num++
		}
	}
	return num
}

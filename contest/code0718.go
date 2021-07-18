package contest

// 5161. 可以输入的最大单词数
func canBeTypedWords(text string, brokenLetters string) int {
	broken := make([]bool, 26)
	for _, v := range brokenLetters {
		broken[v-'a'] = true
	}

	text += " "
	succNum := 0
	curWordFlag := true
	for _, v := range text {
		if v == ' ' {
			if curWordFlag {
				succNum++
			}
			curWordFlag = true
			continue
		}
		if broken[v-'a'] {
			curWordFlag = false
		}
	}
	return succNum
}

// 5814. 新增的最少台阶数
func addRungs(rungs []int, dist int) int {
	idx, pre, cur := 0, 0, 0
	rungNum := 0
	for idx < len(rungs) {
		cur = rungs[idx]
		rungNum += calcRungs(pre, cur, dist)
		pre = cur
		idx++
	}
	return rungNum
}

func calcRungs(pre, cur, dist int) int {
	if cur-pre <= dist {
		return 0
	}
	res := (cur - pre) / dist
	if (cur-pre)%dist == 0 {
		res--
	}
	return res
}

// 5816. 查询最大基因差 TODO TLE
func maxGeneticDifference(parents []int, queries [][]int) []int {
	type Query struct {
		QueriesIdx int
		Value      int
	}
	queriesMap := make(map[int][]Query, len(queries))
	for k, query := range queries {
		queriesMap[query[0]] = append(queriesMap[query[0]], Query{
			QueriesIdx: k,
			Value:      query[1],
		})
	}

	res := make([]int, len(queries))
	for node, queryList := range queriesMap {
		maxList := make([]int, len(queryList))
		for node != -1 {
			for k, v := range queryList {
				if (v.Value ^ node) > maxList[k] {
					maxList[k] = v.Value ^ node
				}
			}
			node = parents[node]
		}

		for k, query := range queryList {
			res[query.QueriesIdx] = maxList[k]
		}
	}

	return res
}

// 5815. 扣分后的最大得分 TODO TLE
func maxPoints(points [][]int) int64 {
	m := len(points)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			dp[0] = points[0]
			continue
		}
		dp[i] = getMaxPoint(points[i], dp[i-1])
	}

	score := dp[0][0]
	for j := 0; j < len(dp[m-1]); j++ {
		if dp[m-1][j] > score {
			score = dp[m-1][j]
		}
	}
	return int64(score)
}

func getMaxPoint(curVal, preMax []int) []int {
	curMax := make([]int, len(preMax))
	for k1, v1 := range curVal {
		for k2, v2 := range preMax {
			if v2+v1-abs(k1, k2) > curMax[k1] {
				curMax[k1] = v2 + v1 - abs(k1, k2)
			}
		}
	}
	return curMax
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

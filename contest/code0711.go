package contest

//"aabca"
//"abc"
//"bbcbaba"
//"uuuuu"
//"eyiccwp"
//"tlpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp"
// 2.计算回文串个数
func countPalindromicSubsequence(s string) int {
	storeMap := make(map[string]bool, 26*26)
	idxMap := make(map[rune]int, 26)
	charNum := make(map[rune]int, 26)
	for k, v := range s {
		charNum[v]++
		if leftIdx, ok := idxMap[v]; ok { // 出现过
			store(s, leftIdx, k, storeMap)
			idxMap[v] = k
		} else {
			idxMap[v] = k
		}
	}

	// 单独处理 单字母重复三次的回文串
	for char, num := range charNum {
		if num >= 3 {
			storeMap[string([]byte{byte(char), byte(char), byte(char)})] = true
		}
	}
	return len(storeMap)
}

// store 存储字符在[left,right)区间内的回文串
func store(s string, left, right int, store map[string]bool) {
	base := []byte{s[left], '0', s[right]}
	for _, char := range s[left+1 : right] {
		base[1] = byte(char)
		store[string(base)] = true
	}
	return
}

// 4.合并搜索二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	if len(trees) == 1 {
		return trees[0]
	}

	// 记录根节点
	rootMap := make(map[int]*TreeNode, len(trees))
	for _, v := range trees {
		rootMap[v.Val] = v
	}

	mergeNode := make(map[*TreeNode]bool, len(trees))
	for _, v := range trees {
		// 已知每颗二叉树最多3个节点，假设它们是平衡二叉树
		if v.Left != nil && rootMap[v.Left.Val] != nil { // 可以合并左节点
			v.Left = rootMap[v.Left.Val]
			mergeNode[v.Left] = true
			if getMaxVal(v.Left) >= v.Val {
				return nil
			}
		}
		if v.Right != nil && rootMap[v.Right.Val] != nil { // 可以合并右节点
			v.Right = rootMap[v.Right.Val]
			mergeNode[v.Right] = true
			if getMinVal(v.Right) <= v.Val {
				return nil
			}
		}
	}

	if len(mergeNode) != len(trees)-1 {
		return nil
	}

	for _, node := range trees {
		if !mergeNode[node] {
			return node
		}
	}
	return nil
}

func getMinVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	max := node.Val
	for node != nil {
		max = node.Val
		node = node.Left
	}
	return max
}
func getMaxVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	max := node.Val
	for node != nil {
		max = node.Val
		node = node.Right
	}
	return max
}

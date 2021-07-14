package learning

import "sort"

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	return 0
}

// 162. 寻找峰值
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	if n == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}

// 82. 删除排序链表中的重复元素 II
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res, cur, pre := head, head, (*ListNode)(nil)
	for cur != nil {
		if cur.Next == nil {
			break
		}

		duplicate := false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			duplicate = true
		}
		if duplicate {
			if pre == nil {
				res = cur.Next
			} else {
				pre.Next = cur.Next
			}
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return res
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] != 0 {
			return nil
		} else {
			return [][]int{nums}
		}
	}

	res := make([][]int, 0, len(nums))
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		sumK := findSumK(nums[i+1:], -nums[i])
		for _, list := range sumK {
			res = append(res, []int{nums[i], list[0], list[1]})
		}
	}
	return res
}
func findSumK(nums []int, k int) [][]int {
	if len(nums) < 2 {
		return nil
	}
	if len(nums) == 2 {
		if nums[0]+nums[1] == k {
			return [][]int{{nums[0], nums[1]}}
		}
		return nil
	}

	var res [][]int
	left, right := 0, len(nums)-1
	for left < right {
		cur := nums[left] + nums[right]
		if cur == k {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
		} else if cur < k {
			left++
		} else {
			right--
		}
		for left > 0 && left < right && nums[left] == nums[left-1] {
			left++
		}
	}
	return res
}

// 844. 比较含退格的字符串
//"ab#c"
//"ad#c"
//"ab##"
//"c#d#"
//"a##c"
//"#a#c"
//"a#c"
//"b"
//"xywrrmp"
//"xywrrmu#p"
func backspaceCompare(s string, t string) bool {
	a, b := len(s)-1, len(t)-1
	for a >= 0 && b >= 0 {
		if s[a] == '#' {
			a -= 2
			continue
		}
		if s[b] == '#' {
			b -= 2
			continue
		}

		if s[a] != s[b] {
			return false
		}
		a--
		b--
	}

	return a == b
}

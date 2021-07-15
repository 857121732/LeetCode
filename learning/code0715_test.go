package learning

import (
	"fmt"
	"testing"
)

func Test_getInterval(t *testing.T) {
	fmt.Println(getInterval([]int{0, 5}, []int{1, 3}))
	fmt.Println(getInterval([]int{2, 5}, []int{1, 6}))
	fmt.Println(getInterval([]int{0, 5}, []int{0, 5}))
	fmt.Println(getInterval([]int{0, 5}, []int{6, 7}))
	fmt.Println(getInterval([]int{0, 5}, []int{5, 6}))
	fmt.Println(getInterval([]int{0, 5}, []int{-3, -2}))
	fmt.Println(getInterval([]int{0, 5}, []int{1, 6}))
	fmt.Println(getInterval([]int{0, 5}, []int{-3, 3}))
}

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func Test_slipSubarray(t *testing.T) {
	fmt.Println(slipSubarray([]int{10, 5, 2, 6}, 0, 0, 10, 100))
}

func Test_numSubarrayProductLessThanK(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

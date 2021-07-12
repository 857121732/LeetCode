package learning

import (
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	t.Log(search(nums, target))
}

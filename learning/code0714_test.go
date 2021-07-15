package learning

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {
	//res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	res := threeSum([]int{-2, 0, 0, 2, 2})
	t.Log(res)
}

func Test_backspaceCompare(t *testing.T) {
	args := [][]string{
		{
			"ab#c",
			"ad#c",
		}, {
			"ab##",
			"c#d#",
		}, {
			"a##c",
			"#a#c",
		}, {
			"a#c",
			"b",
		}, {
			"xywrrmp",
			"xywrrmu#p",
		},
	}
	for _, arg := range args {
		flag := backspaceCompare(arg[0], arg[1])
		fmt.Println(arg[0], arg[1], flag)
	}
}

func Test_handleBack(t *testing.T) {
	fmt.Println(handleBack("a##b#c", 5))
	fmt.Println(handleBack("a##b#c", 4))
	fmt.Println(handleBack("ab#c", 2))
}

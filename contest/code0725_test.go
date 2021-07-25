package contest

import "testing"

func Test_maxCompatibilitySum(t *testing.T) {
	//student := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}
	//teacher := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}
	//student := [][]int{{0,0},{0,0},{0,0}}
	//teacher := [][]int{{1,1},{1,1},{1,1}}
	//student :=[][]int{{1,1,1},{0,0,1},{0, 0, 1}, {0, 1, 0}}
	//teacher := [][]int{{1, 0, 1}, {0, 1, 1}, {0,1,0},{1,1,0}}
	student := [][]int{{0, 0, 1, 0, 1}, {1, 0, 1, 1, 1}}
	teacher := [][]int{{1, 0, 1, 0, 1}, {1, 0, 1, 1, 0}}
	res := maxCompatibilitySum(student, teacher)
	t.Log(res)
}

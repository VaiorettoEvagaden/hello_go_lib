package solution_test

import (
	"reflect"
	"testing"

	"github.com/VaiorettoEvagaden/hello_go_lib/src/solution"
)

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := solution.TwoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

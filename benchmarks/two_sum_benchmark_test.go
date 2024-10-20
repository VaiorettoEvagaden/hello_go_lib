package solution_test

import (
	"testing"

	"github.com/VaiorettoEvagaden/hello_go_lib/src/solution"
)

// !+bench
func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		solution.TwoSum(nums, target)
	}
}

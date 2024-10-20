package solution_test

import (
	"testing"

	"github.com/VaiorettoEvagaden/hello_go_lib/src/solution"
)

// generateLargeSlice generates a slice of integers from 1 to n
func generateLargeSlice(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return nums
}

// !+bench
func BenchmarkTwoSum2(b *testing.B) {
	nums := generateLargeSlice(10000)
	target := 105
	for i := 0; i < b.N; i++ {
		solution.TwoSum(nums, target)
	}
}

func BenchmarkTwoSumBubble(b *testing.B) {
	nums := generateLargeSlice(10000)
	target := 105
	for i := 0; i < b.N; i++ {
		solution.TwoSumBubble(nums, target)
	}
}

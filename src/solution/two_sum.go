package solution

// Time: O(n*n) = O(n^2)
// Space: O(1)
func TwoSum(nums []int, target int) []int {
	// Time: O(n)
	for i := 0; i < len(nums)-1; i++ {
		// Time: O(n)
		for j := i + 1; j < len(nums); j++ {
			// Time: O(1)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSumBubble(nums []int, target int) []int {
	// Time: O(n*n) = O(n^2)
	for i := 0; i < len(nums)-1; i++ {
		// Time: O(n)
		for j := len(nums) - 1; j > i; j-- {
			// Time: O(1)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

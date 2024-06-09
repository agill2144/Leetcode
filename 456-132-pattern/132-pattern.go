func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// Stack to keep track of potential nums[j] and nums[k]
	stack := []int{}
	// Variable to keep track of nums[k] in 132 pattern
	k := -1 << 63 // This will be the smallest value initially

	for i := n - 1; i >= 0; i-- {
		// If we find nums[i] < nums[k], we have a valid 132 pattern
		if nums[i] < k {
			return true
		}
		// Update k and stack with potential nums[j] values
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			k = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
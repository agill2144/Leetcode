func nextPermutation(nums []int)  {
    n := len(nums)
	if n <= 1 {return}
	breachIdx := -1
	for i := n-2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			breachIdx = i
			break
		}
	}
	if breachIdx != -1 {
		ns := -1 // idx
		for i := n-1; i > breachIdx; i-- {
			if nums[i] > nums[breachIdx] && (ns == -1 || nums[i] < nums[ns]) {
				ns = i
			}
		}
		nums[breachIdx], nums[ns] = nums[ns], nums[breachIdx]
	}
	left := breachIdx+1
	right := n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}    
}
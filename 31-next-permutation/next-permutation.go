func nextPermutation(nums []int)  {
    n := len(nums)
    breachIdx := -1
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            breachIdx = i
            break
        }
    }
    if breachIdx != -1 {
        nextGreater := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if nextGreater == -1 || nums[i] < nums[nextGreater] {
                    nextGreater = i
                }
            }
        }

        nums[breachIdx] , nums[nextGreater] = nums[nextGreater], nums[breachIdx]
    }
    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
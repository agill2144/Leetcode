func nextPermutation(nums []int)  {
    n := len(nums)
    breachIdx := -1
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {breachIdx = i; break}
    }
    if breachIdx != -1 {
        nextGIdx := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if nextGIdx == -1 || nums[i] < nums[nextGIdx] {
                    nextGIdx = i
                }
            }
        }

        nums[breachIdx], nums[nextGIdx] = nums[nextGIdx], nums[breachIdx]
    }

    left := breachIdx+1
    right := n-1
    for left < right {nums[left], nums[right] = nums[right], nums[left]; left++; right--}
    
}
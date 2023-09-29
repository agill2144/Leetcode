func nextPermutation(nums []int)  {
    n := len(nums)
    breachIdx := n-2
    for breachIdx >= 0 && nums[breachIdx] >= nums[breachIdx+1] {
        breachIdx--
    }
    
    if breachIdx != -1 {
        fmt.Println(nums[breachIdx])
        nextSmallestIdx := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if nextSmallestIdx == -1 || nums[i] < nums[nextSmallestIdx] {
                    nextSmallestIdx = i
                }
            }
        }
        nums[breachIdx], nums[nextSmallestIdx] = nums[nextSmallestIdx], nums[breachIdx]
    }
    
    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
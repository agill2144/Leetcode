func nextPermutation(nums []int)  {
    breachIdx := -1
    n := len(nums)
    for i := n-1; i >= 0; i-- {
        curr := nums[i]
        next := curr
        if i+1 < n {next = nums[i+1]}
        if curr < next {breachIdx = i; break}
    }

    if breachIdx != -1 {
        nextGreaterIdx := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] && (nextGreaterIdx == -1 || nums[i] < nums[nextGreaterIdx]) {
                nextGreaterIdx = i
            }
        }
        nums[breachIdx], nums[nextGreaterIdx] = nums[nextGreaterIdx], nums[breachIdx]
    }

    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
    
}
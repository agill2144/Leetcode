func nextPermutation(nums []int)  {
    breachIdx := len(nums)-2
    for breachIdx >= 0 {
        if nums[breachIdx] < nums[breachIdx+1] {
            break
        }
        breachIdx--
    }
    
    if breachIdx >= 0 {
        nextGreaterIdx := -1
        for i := len(nums)-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] && (nextGreaterIdx == -1 || nums[i] < nums[nextGreaterIdx]) {
                nextGreaterIdx=i
            }
        }
        
        nums[breachIdx], nums[nextGreaterIdx] = nums[nextGreaterIdx], nums[breachIdx]
    }
    
    left := breachIdx+1
    right := len(nums)-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
    
}

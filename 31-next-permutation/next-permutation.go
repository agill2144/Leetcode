func nextPermutation(nums []int)  {
    breachIdx := -1
    n := len(nums)
    i := n-2
    for i >= 0 {
        if nums[i] < nums[i+1] {
            breachIdx = i
            break
        }
        i--
    }
    if breachIdx != -1 {
        nextG := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if nextG == -1 || nums[i] < nums[nextG] {
                    nextG = i
                }
            }
        }
        nums[breachIdx], nums[nextG] = nums[nextG], nums[breachIdx]
    }

    left := breachIdx+1
    right := n-1
    for left < right {nums[left], nums[right] = nums[right], nums[left]; left++; right--}
    

}
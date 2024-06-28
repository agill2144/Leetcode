func nextPermutation(nums []int)  {
    breachIdx := -1
    n := len(nums)
    for i := n-2; i >= 0; i-- {
        curr := nums[i]
        next := nums[i+1]
        if curr < next {
            breachIdx = i
            break
        }
    }

    if breachIdx  != -1 {
        nextSmallerIdx := -1
        // left most value on right side of breach val
        breachVal := nums[breachIdx]

        for i := n-1; i > breachIdx; i-- {
            curr := nums[i]
            if curr > breachVal && (nextSmallerIdx == -1 || curr < nums[nextSmallerIdx]) {
                nextSmallerIdx = i
            }
        }

        nums[breachIdx], nums[nextSmallerIdx] = nums[nextSmallerIdx], nums[breachIdx]
    }

    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}

/*

    1,2,3,9,8,7
        b     n

    1,2,7,9,8,3
    

*/
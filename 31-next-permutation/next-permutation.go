func nextPermutation(nums []int)  {
    n := len(nums)
    if n <= 1 {return }

    nextSmallerIdx := -1
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            nextSmallerIdx = i
            break
        }
    }
    if nextSmallerIdx != -1 {
        ns := -1
        for i := n-1; i > nextSmallerIdx; i-- {
            if nums[i] > nums[nextSmallerIdx] {
                if ns == -1 || nums[i] < nums[ns] {
                    ns = i
                }
            }
        }
        nums[nextSmallerIdx], nums[ns] = nums[ns], nums[nextSmallerIdx]
    }
    
    left := nextSmallerIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}

func nextPermutation(nums []int)  {
    breachIdx := -1
    n := len(nums)
    
    // find the first adacent pair from right side where left is smaller than right
    // why smaller than right ?
    // because we want to relace the smallest number from right side with the next biggest number
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            breachIdx = i
            break
        }
    }
    
    // look for the next biggest number ( just right of breachIdx value )
    if breachIdx != -1 {
        ns := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if ns == -1 || nums[i] < nums[ns] {
                    ns = i
                }
            }
        }
        // swap the smallest number from right side ( breachIdx) with the next possible number (ns)
        nums[breachIdx], nums[ns] = nums[ns], nums[breachIdx]
    }
    
    // reverse all elements to the right of breachIdx, so that they are all in increasing order from breachIdx
    // why? because we put the next greater number at breachIdx, and now we want the smallest possible permutation to the right of breachIdx
    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
func searchInsert(nums []int, target int) int {
    if nums == nil || len(nums) == 0 {return 0}
    n := len(nums)
    if target <= nums[0] {return 0}
    if target >= nums[n-1] {
        if nums[n-1]==target {return n-1} 
        return n
    }
    left := 0
    right := n-1
    idx := -1

    // find the next closest value to target
    // i.e leftMost on the right side of target
    // ----------|--------|--------
    //         target.  leftmost on right side of target 
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            // found a value on the right side of target
            // save it as a potential ans
            // but we may have overshot, try to find a even closer value!
            // therefore continue searching on left hand side
            // HOWEVER if mid == target, return its idx position
            if nums[mid] == target {return mid}
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}
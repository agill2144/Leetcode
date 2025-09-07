func searchInsert(nums []int, target int) int {
    n := len(nums)
    if target >= nums[n-1] { 
        if nums[n-1] == target {return n-1} 
        return n
    }
    if target <= nums[0] {return 0}
    idx := -1
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {return mid}
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}
func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {return 0}  
    if target < nums[0] {return 0}
    if target > nums[len(nums)-1] {return len(nums)}
    // find the right most on left side of target
    ans := -1
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            if nums[mid] == target {return mid}
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    // because we looked up the right most val on left side of target
    // then target's insert position is that idx+1
    // therefore the +1 here
    return ans+1
}

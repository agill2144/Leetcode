func isMajorityElement(nums []int, target int) bool {
    left := 0
    right := len(nums)-1
    start := -1
    for left <= right {
        mid := left + (right-left) / 2
        if nums[mid] >= target {
            if nums[mid] == target {start = mid}
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if start == -1 { return false}
    
    left = start
    right = len(nums)-1
    end := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            if nums[mid] == target {         
                end = mid   
            }
            left = mid+1
        } else if nums[mid] > target {
            right = mid-1
        }
    }
    size := end-start+1
    return  size > len(nums)/2
}
func search(nums []int, target int) bool {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return true}
        if nums[left] == nums[mid] && nums[mid] == nums[right] {left++; right--; continue}

        // find sorted half
        // left sorted
        if nums[left] <= nums[mid] {
            if target < nums[mid] && target >= nums[left] {
                right = mid-1
            } else {
                left = mid+1
            }
        } else {
            if target > nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return false
}
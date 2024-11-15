func search(nums []int, target int) int {
    n := len(nums)
    left := 0;
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return mid}
        
        // find the sorted half
        // left sorted compared with mid
        if nums[left] <= nums[mid] {
            // check if target is in sorted half
            if target >= nums[left] && target <= nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            }
        } else {
            if target >= nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return -1
}
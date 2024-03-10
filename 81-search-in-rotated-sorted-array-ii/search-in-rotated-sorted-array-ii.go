func search(nums []int, target int) bool {
    left := 0
    right := len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return true}

        // when we cannot guarantee a side to search on
        // that is; left == mid == right
        // shrink our window, try again
        if nums[left] == nums[mid] && nums[mid] == nums[right]{
            left++; right--
            continue
        }


        // original implementation of search in rotated sorted array 1
        // if left is sorted
        if nums[left] <= nums[mid] {
            // target is within left side ?
            if target >= nums[left] && target <= nums[mid] {
                // search this side
                right = mid-1
            } else {
                // search other side
                left = mid+1
            }
        } else { // right side is sorted
            // target within right side?
            if target >= nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
            
        }
    }
    return false
}
func search(nums []int, target int) bool {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] == target {return true}
        
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            // when we cannot determine a sorted half, shrink our window! FUCKING SHIT
            left++
            right--
            continue
        }
        
        // now classic ; search in rotated sorted array by finding the sorted half
        if nums[left] <= nums[mid] {
            if target >= nums[left] && target <= nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            } 
        } else {
            if target <= nums[right] && target >= nums[mid] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return false
}
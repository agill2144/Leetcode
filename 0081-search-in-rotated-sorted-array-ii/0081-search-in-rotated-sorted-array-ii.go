func search(nums []int, target int) bool {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] == target {return true}
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }
        if nums[left] <= nums[mid] {
            if target >= nums[left] && target <= nums[mid] {
                right = mid-1
            }  else {
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
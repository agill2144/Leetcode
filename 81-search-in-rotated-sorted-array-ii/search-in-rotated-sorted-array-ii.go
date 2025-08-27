func search(nums []int, k int) bool {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == k {return true}
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }
        if nums[left] <= nums[mid] {
            if k <= nums[mid] && k >= nums[left] {
                right = mid-1
            } else {
                left = mid+1
            }
        } else if nums[mid] <= nums[right] {
            if k >= nums[mid] && k <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return false
}
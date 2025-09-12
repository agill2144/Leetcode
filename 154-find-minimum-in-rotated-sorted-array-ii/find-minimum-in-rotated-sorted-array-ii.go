func findMin(nums []int) int {
    ans := math.MaxInt64
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == nums[left] && nums[right] == nums[mid] {
            ans = min(ans, nums[mid])
            left++
            right--
            continue
        }

        if nums[left] <= nums[mid] {
            // left sorted
            ans = min(ans, nums[left])
            left = mid+1
        } else {
            // right sorted
            ans = min(ans, nums[mid])
            right = mid-1
        }
    }
    return ans
}
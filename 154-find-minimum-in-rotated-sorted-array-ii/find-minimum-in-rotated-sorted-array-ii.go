func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    ans := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            ans = min(ans, nums[mid])
            left++
            right--
            continue
        }
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            ans = min(ans, nums[mid])
        } 

        if nums[mid] <= nums[right] {
            ans = min(ans, nums[mid])
            right = mid-1
        } else if nums[mid] >= nums[left] {
            ans = min(ans, nums[left])
            left = mid+1
        }
    }
    return ans
}
func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    ans := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == nums[right] && nums[right] == nums[left] {
            ans = min(nums[mid], ans)
            left++
            right--
        } else if nums[left] <= nums[mid] {
            ans = min(ans, nums[left])
            left = mid+1
        } else  {
            ans = min(ans, nums[mid])
            right = mid-1
        }
    }
    return ans
}
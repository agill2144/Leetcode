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
        } else if nums[mid] <= nums[right] {
            ans = min(nums[mid], ans)
            right = mid-1
        } else  {
            left = mid+1
        }
    }
    return ans
}
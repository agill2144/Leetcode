func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    minVal := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            minVal = min(minVal, nums[left])
            left++
            right--
        } else if nums[left] <= nums[mid] { // left sorted -> search right
            minVal = min(minVal, nums[left])
            left = mid+1
        } else if nums[mid] <= nums[right] { // right sorted -> search left
            minVal = min(minVal, nums[mid])
            right = mid-1
        }
    }
    return minVal
}
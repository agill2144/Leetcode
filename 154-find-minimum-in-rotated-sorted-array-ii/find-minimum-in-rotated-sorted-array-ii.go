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

        // mid could an answer
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            ans = min(ans, nums[mid])
        } 
        // when right sorted
        // right is larger than mid
        // therefore again check if mid is a ans
        if nums[mid] <= nums[right] {
            ans = min(ans, nums[mid])
            right = mid-1

        // when left sorted
        // left is smaller than mid
        // therefore check if left is a ans
        } else if nums[mid] >= nums[left] {
            ans = min(ans, nums[left])
            left = mid+1
        }
    }
    return ans
}
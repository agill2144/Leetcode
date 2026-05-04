func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        // we are looking at a sorted win, min is always nums[left]
        if nums[left] < nums[mid] && nums[mid] < nums[right] {return nums[left]}

        // can mid be the min? check n exit early if it is
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        }

        // otherwise min is always on the unsorted half compared to mid
        if nums[left] <= nums[mid] {
            // left sorted
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
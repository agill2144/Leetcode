func peakIndexInMountainArray(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left+(right-left)/2
        if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == n-1  || nums[mid] > nums[mid+1]) {
            return mid
        } else if nums[mid+1] > nums[mid] {
            left = mid+1
        } else {
            right=mid-1
        }
    }
    return -1
}
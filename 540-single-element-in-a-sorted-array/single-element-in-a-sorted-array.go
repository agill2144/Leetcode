func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2

        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == n-1 || nums[mid] != nums[mid+1]) {
            return nums[mid]
        }

        if mid % 2 == 0 {
            if (mid == n-1 || nums[mid] != nums[mid+1]) {
                right = mid-1
            } else {
                left = mid+1
            }
        } else if mid % 2 != 0 {
            if (mid == 0 || nums[mid] != nums[mid-1]) {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return -1
}
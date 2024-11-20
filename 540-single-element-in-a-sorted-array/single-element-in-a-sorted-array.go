func singleNonDuplicate(nums []int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == len(nums)-1 || nums[mid] != nums[mid+1]) {
            return nums[mid]
        }
        if mid % 2 == 0 {
            if mid == len(nums)-1 || nums[mid] == nums[mid+1] {
                left = mid+1
            } else {
                right = mid-1
            }
        } else {
            if (mid == 0 || nums[mid] == nums[mid-1]) {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return -1
}
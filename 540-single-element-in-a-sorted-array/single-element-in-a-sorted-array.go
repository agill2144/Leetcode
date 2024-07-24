func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        even := mid % 2 == 0

        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == n-1 || nums[mid] != nums[mid+1]) {
            return nums[mid]
        }

        // when is mid our answer?
        if even {
            if (mid == n-1 || nums[mid] == nums[mid+1]) {
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
    return ans
}
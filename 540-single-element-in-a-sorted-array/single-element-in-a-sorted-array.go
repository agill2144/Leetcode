func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == n-1 || nums[mid] != nums[mid+1]) {
            return nums[mid]
        }

        if (mid % 2 == 0 && (mid == n-1 || nums[mid] == nums[mid+1])) || (mid % 2 != 0 && (mid == 0 || nums[mid] == nums[mid-1])){
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
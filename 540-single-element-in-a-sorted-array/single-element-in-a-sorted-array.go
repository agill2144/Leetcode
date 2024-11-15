func singleNonDuplicate(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] != nums[mid-1]) && (mid == n-1 || nums[mid] != nums[mid+1]) {return nums[mid]}
        if mid % 2 == 0 {
            // mid = even idx
            if (mid == n-1 || nums[mid] == nums[mid+1]) {
                left = mid+1
            } else {
                right = mid-1
            }
        } else {
            // mid = odd idx
            if (mid == 0 || nums[mid] == nums[mid-1]) {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return -1
}

/*
     0 1 2 3 4  5   6 
    [3,3,7,7,10,10,11]

*/
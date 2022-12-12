/*
    peak element is an element greater than its left and right neighbors
    
    approach: Binary search
    - if mid > mid-1 and mid+1
    - then return mid
    - otherwise binary search on the greater element side
    
    time: o(logn)
    space: o(1)
*/

func findPeakElement(nums []int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right{
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
            return mid
        } else if nums[mid+1] > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    
    return -1
}
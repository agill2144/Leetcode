/*
    The main gist
    - min element is always in unsorted half of the array

    approach: binary search
    - if left element <= right element, we are looking at a sorted array
        - therefore we return smallest element ( i.e left element )
    - mid is min if its smallest amongst its immediate neighbors
    - find the unsorted half, and take the search there

    time = o(logn)
    space = o(1)
*/
func findMin(nums []int) int {
    left := 0
    right := len(nums)-1

    for left <= right {
        // are we looking at a sorted half ?
        if nums[left] <= nums[right] {
            // then return smallest element ( i.e left element )
            return nums[left]
        }

        mid := left + (right - left) / 2
        // is mid smallest ? 
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == len(nums)-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        }

        // left side is sorted ( when compared to mid )
        if nums[left] <= nums[mid] {
            // which means right is unsorted half
            // since min is always is in unsorted half, take search to right side
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
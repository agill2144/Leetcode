/*
    rotated sorted array pattern #2
    - find the sorted half
    - min always lies in the unsorted half
        - unless curr window is already sorted
*/
func findMin(nums []int) int {
    n := len(nums)
    if n == 1 {return nums[0]}
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[left] < nums[right] {
            return nums[left]
        }
        
        // when is mid our answer ?
        // when mid element is smaller than 2 adj neighbors
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        }

        if nums[left] <= nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
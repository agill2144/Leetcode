// rotated sorted array pattern #1: find the sorted half from mid's perspective
// then check if target is in sorted half ; if yes take the search to that side
// if no; take the search to the other side
func search(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }
        
        // is left sorted?
        if nums[mid] >= nums[left] {
            // left sorted; is target in left side ?
            if target >= nums[left] && target <= nums[mid] {
                // yes it is, take search to left
                right = mid-1
            } else {
                // no its not, take search to right
                left = mid+1
            }
        } else {
            // right sorted; is target in right side?
            if target >= nums[mid] && target <= nums[right] {
                // yes it is, take the search to right side
                left = mid+1
            } else {
                // no its not, take the search on left side
                right = mid-1
            }
        }
    }
    return -1
}
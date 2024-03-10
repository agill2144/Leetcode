/*
    The main gist:
    - find sorted half, check if target is within sorted half

    approach: binary search
    - if mid is our target, return
    - find sorted half
    - if target is within sorted half, search in that area
    - otherwise search in the other half

    time = o(logn)
    space = o(1)
*/
func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return mid}

        // if left is sorted
        if nums[left] <= nums[mid] {
            // target is within left side ?
            if target >= nums[left] && target <= nums[mid] {
                // search this side
                right = mid-1
            } else {
                // search other side
                left = mid+1
            }
        } else { // right side is sorted
            // target within right side?
            if target >= nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
            
        }
    }
    return -1
}
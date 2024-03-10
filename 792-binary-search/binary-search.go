// classic binary search
// time = o(logn)
// space = o(1)
func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
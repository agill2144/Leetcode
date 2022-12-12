// min is always on the unsorted half
// unsorted half is half of array that is  unsorted from mid's perspective
// min is an element that is smallest among its immediate neighbors ( left and right )
// so to find mid, find sorted side from mid's perspective - take the search on unsorted half
// at some point, the search on unsorted half will be the smallest sorted half, when this happens return left

func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    if len(nums) == 1 {
        return nums[0]
    }
    for left <= right {
        mid := left + (right-left)/2
        if nums[left] <= nums[right] {
            return nums[left]
        }
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == len(nums)-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        } else if nums[left] <= nums[mid] {
            // left sorted
            left = mid+1
        } else {
            // right sorted
            right = mid-1
        }
    }
    return -1
}
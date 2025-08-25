func search(nums []int, target int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return mid}
        if target > nums[mid] {
            left = mid+1
        }  else {
            right = mid-1
        }
    }
    return -1
}
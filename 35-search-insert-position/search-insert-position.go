// find and save; left most on right side of target
// |----|--|--|--|---------|------------|----|
//               target   leftMost
func searchInsert(nums []int, target int) int {
    n := len(nums)
    if n == 0 {return 0}
    if nums[len(nums)-1] < target {return n}
    left := 0
    right := n-1
    leftMostIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {return mid}
            leftMostIdx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return leftMostIdx
}
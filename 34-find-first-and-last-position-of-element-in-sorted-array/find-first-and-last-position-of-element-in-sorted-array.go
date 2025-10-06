func searchRange(nums []int, target int) []int {
    n := len(nums)
    left := 0
    right := n-1
    leftIdx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            leftIdx = mid
            right = mid-1
        } else if target < nums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if leftIdx == -1 {return []int{-1,-1}}
    left = leftIdx
    right = n-1
    rightIdx := leftIdx
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            rightIdx = mid
            left = mid+1
        } else if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return []int{leftIdx, rightIdx}
}

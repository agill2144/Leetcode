func searchRange(nums []int, target int) []int {
    leftIdx := -1
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left+(right-left)/2
        if target <= nums[mid] {
            if nums[mid] == target {
                leftIdx = mid
            }
            right=mid-1
        } else {
            left = mid+1
        }
    }
    if leftIdx == -1 {
        return []int{-1,-1}
    }
    left = leftIdx
    right = len(nums)-1
    rightIdx := -1
    for left <= right {
        mid := left+(right-left)/2
        if target >= nums[mid] {
            if nums[mid]==target {
                rightIdx = mid
            }
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return []int{leftIdx, rightIdx}
}
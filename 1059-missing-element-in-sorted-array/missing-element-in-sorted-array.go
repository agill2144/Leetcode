func missingElement(nums []int, k int) int {
    offSet := nums[0]
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        current := nums[mid]
        correct := mid+offSet
        missing := current-correct
        if missing < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    rightVal := nums[right]
    correctRightVal := right+offSet
    missingOnLeft := rightVal - correctRightVal
    return rightVal + (k-missingOnLeft)
}
func sortedSquares(nums []int) []int {
    out := make([]int, len(nums))
    left := 0
    right := len(nums)-1
    ptr := len(out)-1
    for left <= right {
        leftSq := nums[left] * nums[left]
        rightSq := nums[right] * nums[right]
        if leftSq > rightSq {
            out[ptr] = leftSq
            left++
        } else {
            out[ptr] = rightSq
            right--
        }
        ptr--
    }
    return out
}
func sortedSquares(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    left := 0
    right := n-1
    ptr := n-1
    for left <= right {
        leftSq := nums[left]*nums[left]
        rightSq := nums[right]*nums[right]
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
func sortedSquares(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    ptr := n-1
    left := 0
    right := n-1
    for left <= right {
        leftVal := nums[left]*nums[left]
        rightVal := nums[right]*nums[right]
        if leftVal > rightVal {
            out[ptr] = leftVal
            left++
        } else {
            out[ptr] = rightVal
            right--
        }
        ptr--
    }
    return out
}
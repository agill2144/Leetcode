func productExceptSelf(nums []int) []int {
    n := len(nums)
    left := make([]int, n)
    left[0] = 1
    for i := 1; i < n; i++ {
        leftRp := left[i-1]
        prev := nums[i-1]
        left[i] = prev * leftRp
    }
    right := make([]int, n)
    right[n-1] = 1
    for i := n-2; i >= 0; i-- {
        rightRp := right[i+1]
        prev := nums[i+1]
        right[i] = rightRp * prev
    }
    out := make([]int, n)
    for i := 0; i < n; i++ {
        out[i] = left[i]*right[i]
    }
    return out
}
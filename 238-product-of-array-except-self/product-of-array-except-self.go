func productExceptSelf(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    out[0] = 1
    for i := 1; i < n; i++ {
        out[i] = nums[i-1]*out[i-1]
    }
    rrp := 1
    for i := n-1; i >= 0; i-- {
        out[i] *= rrp
        rrp *= nums[i]
    }
    return out
}
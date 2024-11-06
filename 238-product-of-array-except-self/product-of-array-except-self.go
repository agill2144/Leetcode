func productExceptSelf(nums []int) []int {
    n := len(nums)
    left := make([]int, n)
    left[0] = 1
    for i := 1; i < n; i++ {
        leftRp := left[i-1]
        prev := nums[i-1]
        left[i] = prev * leftRp
    }
    out := make([]int, n)
    rrp := 1
    out[n-1] = left[n-1]
    for i := n-2; i >= 0; i-- {
        prev := nums[i+1]
        rrp *= prev
        out[i] = rrp*left[i]
    }
    
    return out
}
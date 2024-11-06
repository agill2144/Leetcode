func productExceptSelf(nums []int) []int {
    n := len(nums)
    left := make([]int, n)
    left[0] = 1
    for i := 1; i < n; i++ {
        leftRp := left[i-1]
        prev := nums[i-1]
        left[i] = prev * leftRp
    }
    rrp := 1
    for i := n-2; i >= 0; i-- {
        prev := nums[i+1]
        rrp *= prev
        left[i] *= rrp
    }
    
    return left
}
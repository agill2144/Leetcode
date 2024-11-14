func maxProduct(nums []int) int {
    ans := math.MinInt64
    left := 1
    right := 1
    n := len(nums)
    for i := 0; i < n; i++ {
        left *= nums[i]
        right *= nums[n-1-i]
        ans = max(ans, max(left, right))
        if left == 0 {left = 1}
        if right == 0 {right = 1}
    }
    return ans
}
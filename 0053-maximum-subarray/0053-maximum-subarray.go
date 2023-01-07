func maxSubArray(nums []int) int {
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    out := dp[0]    
    for i := 1; i < len(nums); i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        out = max(out, dp[i])
    }

    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
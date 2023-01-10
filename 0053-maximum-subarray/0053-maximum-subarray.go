func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return -1}
    if len(nums) == 1 {return nums[0]}
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = max(nums[1], nums[1]+nums[0])
    out := max(dp[0], dp[1])

    for i := 2; i < len(nums); i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        out = max(out, dp[i])
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
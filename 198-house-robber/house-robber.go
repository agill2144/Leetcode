func rob(nums []int) int {
    n := len(nums)
    if n <= 1 {
        if n == 1 {return nums[0]}
        return 0
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }
    return dp[n-1]
}
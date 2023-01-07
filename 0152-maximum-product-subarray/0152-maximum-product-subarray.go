func maxProduct(nums []int) int {
    if nums == nil || len(nums) == 0 {return 999999999}
    
    dp := make([][]int, len(nums))
    for i := 0; i < len(dp); i++ {dp[i] = make([]int,2)}
    dp[0][0], dp[0][1] = nums[0], nums[0]    
    maxProd := nums[0]
    for i := 1; i < len(dp); i++ {
        dp[i][0] = min(nums[i], min(nums[i]*dp[i-1][0], nums[i]*dp[i-1][1]))
        dp[i][1] = max(nums[i], max(nums[i]*dp[i-1][0], nums[i]*dp[i-1][1]))
        maxProd = max(maxProd, max(dp[i][0], dp[i][1]))
    }
    return maxProd
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
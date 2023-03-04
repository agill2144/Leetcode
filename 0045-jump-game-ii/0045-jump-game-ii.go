func jump(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp)-1; i++ {
        dp[i] = 10001
    }
    
    for i := len(nums)-2; i>=0;  i-- {
        jumpDist := nums[i]
        if jumpDist == 0 {continue}
        for j := jumpDist; j >= 1; j-- {
            if i+j < len(nums) {
                dp[i] = min(dp[i], 1+dp[i+j])
            }
        }
    }
    return dp[0]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}
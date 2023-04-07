func canJump(nums []int) bool {
    dp := make([]bool, len(nums))
    dp[len(nums)-1] = true
    
    for i := len(nums)-2; i>=0; i-- {
        jumpDist := nums[i]
        for j := 1; j <= jumpDist; j++ {
            if i+j < len(nums) && dp[i+j] {
                dp[i] = true
                break
            }
        }
    }
    return dp[0]
}
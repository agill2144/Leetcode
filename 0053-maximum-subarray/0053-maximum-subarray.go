func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    maxTotal := nums[0]
    
    for i := 1; i < len(nums); i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        if dp[i] > maxTotal {maxTotal = dp[i]}
    }
    return maxTotal    
}

func max(x, y int) int {
    if x > y {return x} 
    return y
}
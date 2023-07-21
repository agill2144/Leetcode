func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := 0; i < n; i++ {dp[i] = 1}

    maxLen := 1
    for i := 1; i < n; i++{
        curr := nums[i]
        for j := i-1; j >= 0; j-- {
            if curr > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
            maxLen = max(dp[i], maxLen)
        }
    }
    return maxLen
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
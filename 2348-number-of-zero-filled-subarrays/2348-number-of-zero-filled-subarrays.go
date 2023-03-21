func zeroFilledSubarray(nums []int) int64 {
    dp := make([]int, len(nums))
    var total int64 = 0
    if nums[0] == 0 {
        dp[0] = 1
        total = 1
    }
    for i := 1; i < len(nums); i++ {
        if nums[i] == 0 {
            dp[i]++
            if nums[i-1] == 0 {
                dp[i] += dp[i-1]
            }
        }
        total += int64(dp[i])
    }
    // fmt.Println(dp)
    return total
}
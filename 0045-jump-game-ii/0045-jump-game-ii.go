func jump(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++{dp[i] = math.MaxInt64-100}
    dp[len(dp)-1] = 0

    for i := len(nums)-1; i >= 0; i-- {
        numJumps := nums[i]
        for j := numJumps; j >= 1; j-- {
            if j+i < len(nums) {
                dp[i] = min(dp[i], dp[i+j]+1)
            }
        }
    }
    return dp[0]
}
func min(x, y int) int {
    if x < y {return x}
    return y
}
/*
    bottom up dp
    time: o(n) + o(n*j)
    space: o(n)
*/
func jump(nums []int) int {
    dp := make([]int, len(nums))
    
    // fill it with Inf value
    for i := 0; i < len(dp); i++ {dp[i] = math.MaxInt64-100}
    dp[len(dp)-1] = 0
    
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
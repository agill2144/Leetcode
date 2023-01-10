func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return -1}
    if len(nums) == 1 {return nums[0]}
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    out := dp[0]

    for i := 1; i < len(nums); i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        out = max(out, dp[i])
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}



/*
    approach: brute force
    - nested loop
    - keep track of running sum
    - save running into max whenever needed
    
    time: o(n^2)
    space: o(1)
    

func maxSubArray(nums []int) int {
    max := math.MinInt64
    for i := 0; i < len(nums); i++ {
        running := 0
        for j := i; j < len(nums); j++ {
            running += nums[j]
            if running > max {
                max = running
            }
        }
    }
    return max
}
*/
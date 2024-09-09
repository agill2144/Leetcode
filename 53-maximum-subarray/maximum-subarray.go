func maxSubArray(nums []int) int {
    // kadanes greedy algo, try to evaluate each idx as a new starting point
    // start a new eval if rSum < 0 
    rSum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        maxSum = max(maxSum, rSum)
        if rSum < 0 {
            rSum = 0
        }
    }
    return maxSum
}
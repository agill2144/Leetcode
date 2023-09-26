func maxSubArray(nums []int) int {
    runningSum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        runningSum += nums[i]
        if runningSum > maxSum {
            maxSum = runningSum
        }
        if runningSum < 0 {
            runningSum = 0
        }
    }
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
func maxSubArray(nums []int) int {
    runningSum := 0
    runningStart := 0
    maxSum := math.MinInt64
    start := -1
    end := -1
    for i := 0; i < len(nums); i++ {
        runningSum += nums[i]
        if runningSum > maxSum {
            maxSum = runningSum
            start = runningStart
            end = i
        }
        if runningSum < 0 {
            runningSum = 0
            runningStart = i+1
        }
    }
    // throwing values away so go-complier does not complain of unused values
    // we could print it too
    _ = start
    _ = end
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
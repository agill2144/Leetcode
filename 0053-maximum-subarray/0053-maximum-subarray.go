func maxSubArray(nums []int) int {
    runningSum := 0
    runningStart := 0
    maxSum := nums[0]
    start := 0
    end := 0
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
    _ = start
    _ = end
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
func maximumSumScore(nums []int) int64 {
    // 2 pass
    // 1. get the total sum
    // 2. use the total sum to apply the 2 rules
    var total int64
    for i := 0; i < len(nums); i++ {total += int64(nums[i])}
    var maxScore int64 = math.MinInt64 
    var runningSum int64
    
    for i := 0; i < len(nums); i++ {
        runningSum += int64(nums[i])
        remainingRightSum := total-runningSum + int64(nums[i])
        maxScore = max(maxScore, max(runningSum, remainingRightSum))

    }
    return maxScore
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}
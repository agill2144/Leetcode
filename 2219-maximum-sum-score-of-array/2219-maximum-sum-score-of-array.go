func maximumSumScore(nums []int) int64 {
    var total  int64 = 0
    for i := 0; i < len(nums); i++ {
        total += int64(nums[i])
    }
    var maxScore int64 = math.MinInt64
    var rSum int64 = 0
    for i := 0; i < len(nums); i++ {
        rSum += int64(nums[i])
        right := (total-rSum) + int64(nums[i])
        maxScore = max(maxScore, max(rSum, right))
    }
    return maxScore
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}
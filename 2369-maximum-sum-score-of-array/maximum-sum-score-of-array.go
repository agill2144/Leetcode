func maximumSumScore(nums []int) int64 {
    var ans int64 = math.MinInt64
    var total int64
    for i := 0; i < len(nums); i++ {total += int64(nums[i])}
    var rSum int64
    for i := 0; i < len(nums); i++ {
        rSum += int64(nums[i])
        rightSum := total - rSum + int64(nums[i])
        ans = max(ans, max(rSum, rightSum))
    }
    return ans
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}

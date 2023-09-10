func maximumSumScore(nums []int) int64 {
    var total int64
    for i := 0; i < len(nums); i++ {
        total += int64(nums[i])
    }
    var ans int64 = math.MinInt64
    var leftRunning int64
    for i := 0; i < len(nums); i++ {
        leftRunning += int64(nums[i])
        var right int64 = (total - leftRunning) + int64(nums[i])
        ans = max(ans, max(leftRunning, right))
    }
    return ans
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}
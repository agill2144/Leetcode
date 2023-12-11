func maximumSumScore(nums []int) int64 {
    var total int64
    var ans int64 = math.MinInt64
    for i := 0; i < len(nums); i++ {
        total += int64(nums[i])
    }
    var running int64
    for i := 0; i < len(nums); i++ {
        running += int64(nums[i])
        remaining := (total-running) + int64(nums[i])
        ans = max(ans, max(running, remaining))
    }
    return ans
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}
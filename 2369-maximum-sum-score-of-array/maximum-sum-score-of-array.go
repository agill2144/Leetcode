func maximumSumScore(nums []int) int64 {
    var total int64 = 0
    for i := 0; i < len(nums); i++ {
        total += int64(nums[i])
    }
    var maxSum int64 = math.MinInt64
    var rSum int64 = 0
    for i := 0; i < len(nums); i++ {
        rSum += int64(nums[i])
        var rightSum int64 = total-rSum + int64(nums[i])
        maxSum = max(maxSum, max(rSum, rightSum))
    }
    return maxSum
}
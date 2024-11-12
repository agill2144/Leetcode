func maxSubArray(nums []int) int {
    maxSum := math.MinInt64
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        maxSum = max(maxSum, rSum)
        if rSum < 0 {rSum = 0}
    }
    return maxSum
}
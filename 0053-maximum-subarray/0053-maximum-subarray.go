func maxSubArray(nums []int) int {
    sum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum > maxSum {
            maxSum = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return maxSum
}
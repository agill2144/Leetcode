func maxSubArray(nums []int) int {
    max := math.MinInt64
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum > max {
            max = sum
        }
        if sum < 0 {sum = 0}
    }
    return max
}
func maxSubArray(nums []int) int {
    rs := nums[0]
    maxSum := nums[0]
    for i := 1; i < len(nums); i++ {
        rs += nums[i]
        rs = max(rs, nums[i])
        maxSum = max(rs, maxSum)
    }
    return maxSum
}
func max(x, y int) int {
    if x > y {return x}
    return y
}
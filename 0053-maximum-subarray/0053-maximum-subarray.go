func maxSubArray(nums []int) int {
    sum := 0
    maxSum := nums[0]
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        maxSum = max(sum, maxSum)
        if sum < 0 {sum = 0}
    }
    return maxSum
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
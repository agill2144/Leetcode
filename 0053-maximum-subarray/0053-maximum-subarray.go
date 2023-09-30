func maxSubArray(nums []int) int {
    sum := 0
    maxSum := math.MinInt64
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum > maxSum {maxSum = sum}
        
        // is it worth keeping this rolling sum / subarray ?
        // its not worth keeping if the sum is negative 
        // its worth keeping if the sum in positive
        if sum < 0 {
            sum = 0
        }
    }
    return maxSum
}
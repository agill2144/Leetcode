func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := n*(n+1)/2 
    sum := 0
    for i := 0; i < n; i++ {sum += nums[i]}
    return expectedSum - sum
}
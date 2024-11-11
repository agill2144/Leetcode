func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := n*(n+1)/2
    actualSum := 0
    for i := 0; i < len(nums); i++ {actualSum += nums[i]}
    return expectedSum - actualSum    
}
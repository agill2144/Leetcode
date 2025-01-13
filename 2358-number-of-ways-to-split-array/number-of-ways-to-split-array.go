func waysToSplitArray(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }
    count := 0
    leftSum := 0
    for i := 0; i < len(nums)-1; i++ {
        leftSum += nums[i]
        rightSum := total-leftSum
        if leftSum >= rightSum {count++}
    }
    return count
}
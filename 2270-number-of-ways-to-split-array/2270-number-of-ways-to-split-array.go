func waysToSplitArray(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }

    count := 0
    runningLeftSum := 0
    for i := 0; i < len(nums)-1; i++ {
        runningLeftSum += nums[i]
        rightSum := total-runningLeftSum
        if runningLeftSum >= rightSum {
            count++
        }
    }
    return count
}

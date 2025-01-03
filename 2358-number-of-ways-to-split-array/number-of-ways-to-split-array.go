func waysToSplitArray(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {total += nums[i]}
    sum := 0
    count := 0
    for i := 0; i < len(nums)-1; i++ {
        sum += nums[i]
        rightSum := total-sum
        if sum >= rightSum {count++}
    }
    return count
}
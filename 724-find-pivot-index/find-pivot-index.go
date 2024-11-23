func pivotIndex(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }
    rSum := 0
    for i := 0; i < len(nums); i++ {
        leftSum := rSum
        rightSum := total - leftSum - nums[i]
        if leftSum == rightSum {return i}
        rSum += nums[i]
    }
    return -1
}
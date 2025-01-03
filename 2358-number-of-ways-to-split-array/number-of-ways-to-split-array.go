func waysToSplitArray(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++ {total += nums[i]}
    sum := 0
    count := 0
    // why upto n-1 ? 
    // because of "There is at least one element to the right of i"
    // meaning, we must have 1 num in the right subarray/ right prefix sum
    for i := 0; i < len(nums)-1; i++ {
        sum += nums[i]
        rightSum := total-sum
        if sum >= rightSum {count++}
    }
    return count
}
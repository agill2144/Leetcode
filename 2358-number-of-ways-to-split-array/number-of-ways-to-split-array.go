func waysToSplitArray(nums []int) int {
    n := len(nums)
    total := 0
    for i := 0; i < n; i++ {total += nums[i]}
    left := 0
    count := 0
    for i := 0; i < n-1; i++ {
        left += nums[i]
        right := total-left
        if left >= right {
            count++
        }
    }
    return count
}
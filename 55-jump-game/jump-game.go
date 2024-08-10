func canJump(nums []int) bool {
    maxIdx := 0
    i := 0
    for i <= maxIdx && i < len(nums) {
        maxIdx = max(maxIdx, i+nums[i])
        i++
    }
    return maxIdx >= len(nums)-1
}
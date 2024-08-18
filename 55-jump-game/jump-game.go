func canJump(nums []int) bool {
    farthestIdx := 0
    i := 0
    n := len(nums)
    for i < n && i <= farthestIdx {
        farthestIdx = max(farthestIdx, i+nums[i])
        i++
    }
    return farthestIdx >= n-1
}
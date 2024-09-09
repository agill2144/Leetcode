func canJump(nums []int) bool {
    farthestIdx := 0
    i := 0
    for i <= farthestIdx && i < len(nums) {
        farthestIdx = max(farthestIdx, nums[i]+i)
        if farthestIdx >= len(nums)-1 {return true}
        i++
    }
    return false
}
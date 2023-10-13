func canJump(nums []int) bool {
    targetIdx := len(nums)-1
    for i := len(nums)-2; i >= 0; i-- {
        if i + nums[i] >= targetIdx {
            targetIdx = i
        }
    }
    return targetIdx == 0
}
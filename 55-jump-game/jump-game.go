func canJump(nums []int) bool {
    n := len(nums)
    jumpToIdx := n-1
    for i := n-2; i >= 0; i-- {
        if i + (nums[i]) >= jumpToIdx {
            jumpToIdx = i
        }
    }
    return jumpToIdx == 0
}
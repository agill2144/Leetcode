func canJump(nums []int) bool {
    n := len(nums)
    target := n-1    
    for i := n-2; i>=0; i-- {
        jumpDist := nums[i]
        if i+jumpDist >= target {
            target = i
        }
    }
    return target == 0
}
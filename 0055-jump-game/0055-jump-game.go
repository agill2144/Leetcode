func canJump(nums []int) bool {
    if len(nums) <= 1 {return true}
    if nums[0] == 0 {return false}
    
    target := len(nums)-1
    for i := len(nums)-2; i >= 0; i-- {
        if i+nums[i] >= target {
            target = i
        }
    }
    return target == 0 
}
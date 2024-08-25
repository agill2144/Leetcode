func jump(nums []int) int {
    left := 0
    right := 0
    jumps := 0
    if len(nums) <= 1 {return 0}
    for right < len(nums)-1 {
        farthest := left
        for i := left; i <= right; i++ {
            farthest = max(farthest, i+nums[i])
        }
        left = right+1
        right = farthest
        jumps++ 
        if right >= len(nums)-1 {return jumps}
    }
    return -1
}
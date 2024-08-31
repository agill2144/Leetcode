func jump(nums []int) int {
    n := len(nums)
    left := 0
    right := 0
    jumps := 0
    for right < n-1 {
        farthestIdx := -1
        for i := left; i <= right; i++ {
            farthestIdx = max(farthestIdx, i+nums[i])
        }
        if farthestIdx <= right {panic("could not jump out of current range/block")}
        left = right+1
        right = farthestIdx
        jumps++
    }
    return jumps
}
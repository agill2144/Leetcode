func jump(nums []int) int {
    left := 0
    right := 0
    n := len(nums)
    if n <= 1 {return 0}
    jumps := 0
    for right < n-1 {
        farthestIdx := -1
        for i := left; i <= right; i++ {
            farthestIdx = max(farthestIdx, i+nums[i])
        }
        left = right+1
        right = farthestIdx
        jumps++
        if right >= n-1 {break}
    }
    return jumps
}
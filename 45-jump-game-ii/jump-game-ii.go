func jump(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}
    left := 0
    right := 0
    jumps := 0
    for left < n && right < n {
        farthest := -1
        for i := left; i <= right; i++ {
            farthest = max(farthest,i+nums[i])
        }
        left = right+1
        right = farthest
        jumps++
        if right >= n-1 {break}
    }
    return jumps
}
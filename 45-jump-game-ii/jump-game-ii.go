func jump(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}
    left := 0
    right := 0
    count := 0
    for right < n-1 {
        farthest := -1
        for i := left; i <= right; i++ {
            farthest = max(farthest, i+nums[i])
        }        
        left = right+1
        right = farthest
        count++
        if right >= n-1 {break}
    }
    return count
}
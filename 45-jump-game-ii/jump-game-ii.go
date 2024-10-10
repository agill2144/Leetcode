func jump(nums []int) int {
    count := 0
    left := 0
    right := 0
    for right < len(nums)-1 {
        nextIdx := right
        for i := left; i <= right; i++ {
            nextIdx = max(nextIdx, nums[i]+i)
        }
        left = right+1
        right = nextIdx
        count++
        if right >= len(nums)-1 {break}
    }
    return count
}
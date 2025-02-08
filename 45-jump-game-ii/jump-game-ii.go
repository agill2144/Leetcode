func jump(nums []int) int {
    if len(nums) <= 1 {return 0}
    windows := 0
    right := 0
    i := 0
    maxRight := 0
    for i < len(nums) {
        maxRight = max(maxRight,i + nums[i])
        if i == right {windows++; right = maxRight}
        if right >= len(nums)-1 {
            break
        }
        i++
    }
    return windows
}
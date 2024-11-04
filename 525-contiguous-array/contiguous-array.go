func findMaxLength(nums []int) int {
    idx := map[int]int{0:-1}
    count := 0
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {count--} else {count++}
        left, ok := idx[count]
        if ok {
            maxLen = max(maxLen, i-left)
        } else {
            idx[count] = i
        }
    }
    return maxLen
}
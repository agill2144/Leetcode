func findMaxLength(nums []int) int {
    count := 0
    maxLen := 0
    idxs := map[int]int{0:-1}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {count++} else {count--}
        lastIdx, ok := idxs[count]
        if ok {
            maxLen = max(maxLen, i-lastIdx)
        } else {
            idxs[count] = i
        }
    }
    return maxLen
}
func findMaxLength(nums []int) int {
    idx := map[int]int{0:-1}
    maxSize := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {count++}
        if nums[i] == 0 {count--}
        val, ok := idx[count]
        if ok {
            maxSize = max(maxSize, i-val)
        }
        if _, ok := idx[count]; !ok {
            idx[count] = i
        }
    }
    return maxSize
}
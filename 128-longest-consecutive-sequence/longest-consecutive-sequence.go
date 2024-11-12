func longestConsecutive(nums []int) int {
    set := map[int]bool{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = true
    }
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        if set[nums[i]-1] {continue}
        k := nums[i]
        size := 0
        for set[k] {
            size++
            k++
        }
        maxLen = max(maxLen, size)
    }
    return maxLen
}
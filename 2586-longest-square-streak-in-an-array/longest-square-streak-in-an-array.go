func longestSquareStreak(nums []int) int {
    set := map[int]bool{}
    smallest := math.MaxInt64
    largest := math.MinInt64
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = true
        smallest = min(smallest, nums[i])
        largest = max(largest, nums[i])
    }
    maxSize := 1
    for i := smallest; i <= largest; i++ {
        count := 0
        val := i
        for set[val] {
            count++
            val *= val
        }
        maxSize = max(maxSize, count)
    }
    if maxSize == 1 {maxSize = -1}
    return maxSize
}
func longestSquareStreak(nums []int) int {
    set := map[int]bool{}
    n := len(nums)
    maxN := math.MinInt64
    for i := 0; i < n; i++ {set[nums[i]] = true; maxN = max(maxN, nums[i])}
    maxSize := 0
    for i := 0; i < n; i++ {
        curr := nums[i]
        start := curr
        size := 1
        for set[start*start] && start*start <= maxN {
            start *= start
            size++
        }
        maxSize = max(maxSize, size)
    }
    if maxSize == 1 {maxSize = -1}
    return maxSize
}


func longestSquareStreak(nums []int) int {
    set := map[int]bool{}
    start := math.MaxInt64
    end := math.MinInt64
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = true
        start = min(start, nums[i])
        end = max(end, nums[i])
    }
    maxStreak := 1
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        count := 0
        for set[curr] {
            curr = curr * curr
            count++
        }
        maxStreak = max(maxStreak, count)
    }
    if maxStreak == 1 {return -1}
    return maxStreak
}
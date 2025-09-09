// hashing: longest consec seq pattern
// tc = o(n) + o(n * log n)
// sc = o(n)
func longestSquareStreak(nums []int) int {
    set := map[float64]bool{}
    for i := 0; i < len(nums); i++ {set[float64(nums[i])] = true}
    maxStreak := 0
    for i := 0; i < len(nums); i++ {
        sqrt := math.Sqrt(float64(nums[i]))
        if set[sqrt] {continue}
        start := float64(nums[i])
        count := 0
        for set[start] {
            count++
            start *= start
        }
        maxStreak = max(maxStreak, count)
    }
    if maxStreak == 1 {maxStreak = -1}
    return maxStreak
}
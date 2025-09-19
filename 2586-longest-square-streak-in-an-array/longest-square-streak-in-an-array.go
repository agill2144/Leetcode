func longestSquareStreak(nums []int) int {
    set := map[float64]bool{}
    for i := 0; i < len(nums); i++ {
        set[float64(nums[i])] = true
    }
    ans := -1
    for i := 0; i < len(nums); i++ {
        sqrt := math.Sqrt(float64(nums[i]))
        if !set[sqrt] {
            count := 0
            start := float64(nums[i])
            for set[start] {
                count++
                start *= start
            }
            if count > 1 {ans = max(ans, count)}
        }
    }
    return ans
}
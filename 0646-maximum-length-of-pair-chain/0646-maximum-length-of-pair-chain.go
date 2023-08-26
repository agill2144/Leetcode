// dp vibes
func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int)bool {
        return pairs[i][0] < pairs[j][0]
    })
    
    dp := make([]int, len(pairs))
    dp[0] = 1
    someMax := 1
    for i := 1; i < len(pairs); i++ {
        curr := pairs[i]
        dp[i] = 1
        for j := i-1; j >= 0; j-- {
            prev := pairs[j]
            if curr[0] > prev[1] {
                dp[i] = max(dp[i], dp[j]+1)
            }
            someMax = max(someMax, dp[i])
        }
        
    }
    return someMax
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
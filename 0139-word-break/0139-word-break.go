func wordBreak(s string, wordDict []string) bool {
    set := map[string]struct{}{}
    for _, word := range wordDict {
        set[word] = struct{}{}
    }
    
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 1; i < len(dp); i++ {
        for j := 0; j < i; j++ {
            if dp[j]{
                subStr := string(s[j:i])
                if _, ok := set[subStr]; ok {
                    dp[i] = true
                    break
                }
            } 
        }
    }
    return dp[len(dp)-1]
}
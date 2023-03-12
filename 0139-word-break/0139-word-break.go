func wordBreak(s string, wordDict []string) bool {
    set := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ {set[wordDict[i]] = struct{}{}}
    dp := make([]bool, len(s))
    for i := 0; i < len(s); i++ {
        for j := i; j >= 0; j-- {
            if j == 0 || dp[j-1]  {
                if _, ok := set[string(s[j:i+1])]; ok {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[len(dp)-1]
}
func wordBreak(s string, wordDict []string) []string {
    wordSet := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ {
        wordSet[wordDict[i]] = struct{}{}
    }
    out := []string{}
    
    var dfs func(start int, sentence string) 
    dfs = func(start int, sentence string) {
        // base
        if start == len(s) {
            out = append(out, sentence)
            return
        }
        
        // logic
        for i := start ; i < len(s); i++ {
            subStr := s[start:i+1]
            if _, ok := wordSet[subStr]; ok {
                newSent := sentence + " " + subStr
                if sentence == "" {
                    newSent = subStr
                }
                dfs(i+1, newSent)
            }
        }
    }
    dfs(0, "")
    return out
}
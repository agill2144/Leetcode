func wordBreak(s string, wordDict []string) []string {

    set := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ { set[wordDict[i]] = struct{}{} }

    out := []string{}
    var dfs func(start int, sent string)
    dfs = func(start int, sent string) {
        // base
        if start == len(s) {
            
            out = append(out, strings.TrimSpace(sent))
            return
        }
        
        
        // logic
        for i := start; i < len(s); i++ {
            subStr := string(s[start:i+1])
            if _, ok := set[subStr]; ok {
                dfs(i+1, sent+" "+subStr)
            }
        }
    }
    dfs(0, "")
    return out
}
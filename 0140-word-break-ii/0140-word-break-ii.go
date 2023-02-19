func wordBreak(s string, wordDict []string) []string {

    set := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ { set[wordDict[i]] = struct{}{} }

    out := []string{}
    var dfs func(start int, sent []string)
    dfs = func(start int, sent []string) {
        // base
        if start == len(s) {
            out = append(out, strings.Join(sent, " "))
            return
        }
        
        
        // logic
        for i := start; i < len(s); i++ {
            subStr := string(s[start:i+1])
            if _, ok := set[subStr]; ok {
                sent = append(sent, subStr)
                dfs(i+1, sent)
                sent = sent[:len(sent)-1]
            }
        }
    }
    dfs(0, []string{})
    return out
}